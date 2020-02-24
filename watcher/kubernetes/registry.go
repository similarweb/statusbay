package kuberneteswatcher

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"statusbay/watcher/kubernetes/common"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	eventwatch "k8s.io/apimachinery/pkg/watch"
)

const (
	// applyVersionFormat describe the format of apply versions
	applyVersionFormat = "%s-%s-%s-%s"
)

type Resources struct {
	Deployments  map[string]*DeploymentData  `json:"Deployments"`
	Daemonsets   map[string]*DaemonsetData   `json:"Daemonsets"`
	Statefulsets map[string]*StatefulsetData `json:"Statefulsets"`
}

//DBSchema is a struct that save as json in given storage
type DBSchema struct {
	Application           string                             `json:"Application"`
	Cluster               string                             `json:"Cluster"`
	Namespace             string                             `json:"Namespace"`
	CreationTimestamp     int64                              `json:"CreationTimestamp"`
	ReportTo              []string                           `json:"ReportTo"`
	DeployBy              string                             `json:"DeployBy"`
	DeploymentDescription common.DeploymentStatusDescription `json:"DeploymentDescription"`
	Resources             Resources                          `json:"Resources"`
}

// ApplyEvent describe the new Kubernetes apply details for create/skip/delete new application
type ApplyEvent struct {
	Event        string
	ApplyName    string
	ResourceName string
	Namespace    string
	Kind         string
	Hash         uint64
	Annotations  map[string]string
	Labels       map[string]string
}

// RegistryRow defined row data of deployment
type RegistryRow struct {
	applyID                          string
	finish                           bool
	beforeFinish                     bool
	status                           common.DeploymentStatus
	ctx                              context.Context
	cancelFn                         context.CancelFunc
	collectDataAfterDeploymentFinish time.Duration
	DBSchema                         DBSchema
}

// RegistryManager defined multiple rows data
type RegistryManager struct {
	clusterName                 string
	registryData                map[string]*RegistryRow
	saveInterval                time.Duration
	checkFinishDelay            time.Duration
	collectDataAfterApplyFinish time.Duration
	saveLock                    *sync.Mutex
	applyLock                   *sync.Mutex
	storage                     Storage
	reporter                    *ReporterManager
	lastDeploymentHistory       map[string]time.Time
}

// NewRegistryManager create new schema registry instance
func NewRegistryManager(saveInterval time.Duration, checkFinishDelay time.Duration, collectDataAfterApplyFinish time.Duration, storage Storage, reporter *ReporterManager, clusterName string) *RegistryManager {
	if clusterName == "" {
		log.Panic("cluster name is a mandatory field")
		os.Exit(1)
	}

	return &RegistryManager{
		clusterName:                 clusterName,
		saveInterval:                saveInterval,
		checkFinishDelay:            checkFinishDelay,
		collectDataAfterApplyFinish: collectDataAfterApplyFinish,
		storage:                     storage,
		reporter:                    reporter,

		registryData:          make(map[string]*RegistryRow),
		lastDeploymentHistory: make(map[string]time.Time),
		saveLock:              &sync.Mutex{},
		applyLock:             &sync.Mutex{},
	}
}

// Serve will start listening schema registry request
func (dr *RegistryManager) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		for {
			select {
			case <-time.After(dr.saveInterval):
				dr.save()
			case <-ctx.Done():
				log.Warn("registry save schema has been shut down")
				wg.Done()
				return
			}
		}
	}()

}

// LoadRunningApps TODO:: fix me
func (dr *RegistryManager) LoadRunningApplies() []*RegistryRow {

	rows := []*RegistryRow{}
	apps, _ := dr.storage.GetAppliesByStatus(common.ApplyStatusRunning)
	log.WithField("count", len(apps)).Info("loading running job from database")

	for applyID, appSchema := range apps {

		encodedID := generateID(appSchema.Application, appSchema.Namespace, dr.clusterName)
		ctx, cancelFn := context.WithCancel(context.Background())

		row := RegistryRow{
			applyID:  applyID,
			ctx:      ctx,
			cancelFn: cancelFn,
			finish:   false,
			status:   common.ApplyStatusRunning,
			DBSchema: appSchema,
		}
		go row.isFinish(dr.checkFinishDelay)
		dr.registryData[encodedID] = &row

		rows = append(rows, &row)

	}

	return rows

}

func (dr *RegistryManager) NewApplyEvent(data ApplyEvent) *RegistryRow {

	dr.applyLock.Lock()
	defer dr.applyLock.Unlock()

	var appRegistry *RegistryRow

	if data.Event == fmt.Sprintf("%v", eventwatch.Deleted) {
		// Check if the resource already detected in StatusBasy
		appRegistry = dr.Get(data.ApplyName, data.Namespace, data.Event)

		dr.deleteAppliedVersion(data.ResourceName, data.Namespace, data.Kind)
		if appRegistry != nil && appRegistry.beforeFinish {
			return nil
		}

		// Check if the resource in running status, it mean that the resource apply to kubernetes,
		// the apply not finished, and Kubernetes got event for delete this resource
		runningApply := dr.Get(data.ApplyName, data.Namespace, "")

		// if apply found in registry, the apply going to be canceled
		if runningApply != nil {
			lg := runningApply.Log()
			lg.Info("apply was canceled, got delete event")
			go runningApply.Stop(common.ApplyCanceled, common.ApplyStatusDescriptionCanceled)
		}
	} else {

		appRegistry = dr.Get(data.ApplyName, data.Namespace, "")
		if !dr.updateAppliesVersionHistory(data.ResourceName, data.Namespace, data.Kind, data.Hash) {
			log.WithFields(log.Fields{
				"resource_name": data.ResourceName,
				"namespace":     data.Namespace,
				"cluster":       dr.clusterName,
				"resource_kind": data.Kind,
			}).Debug("resource already deployed")
			return nil
		}
	}

	// If apply not found in memory (first event of the apply), we needs to create new application apply
	// for create new strcut that include all the resources
	if appRegistry == nil {
		status := common.ApplyStatusRunning
		if data.Event == fmt.Sprintf("%v", eventwatch.Deleted) {
			status = common.ApplyStatusDeleted
		}

		appRegistry = dr.NewApplication(data.ApplyName, data.Namespace, data.Annotations, status)
	}

	return appRegistry

}

// NewApplication will creates a new deployment row
func (dr *RegistryManager) NewApplication(appName string, namespace string, annotations map[string]string, status common.DeploymentStatus) *RegistryRow {

	encodedID := generateID(appName, namespace, dr.clusterName)
	reportTo := GetMetadataByPrefix(annotations, fmt.Sprintf("%s/%s-", annotationPrefix, annotationPrefixAllReporter))
	deployBy := GetMetadata(annotations, fmt.Sprintf("%s/%s", annotationPrefix, annotationReportDeployBy))
	deployTime := time.Now().Unix()
	ctx, cancelFn := context.WithCancel(context.Background())

	row := RegistryRow{
		applyID:                          "",
		ctx:                              ctx,
		cancelFn:                         cancelFn,
		finish:                           false,
		status:                           status,
		collectDataAfterDeploymentFinish: dr.collectDataAfterApplyFinish,
		DBSchema: DBSchema{
			Application:           appName,
			Cluster:               dr.clusterName,
			Namespace:             namespace,
			CreationTimestamp:     deployTime,
			ReportTo:              reportTo,
			DeployBy:              deployBy,
			DeploymentDescription: common.ApplyStatusDescriptionRunning,
			Resources: Resources{
				Deployments:  make(map[string]*DeploymentData),
				Daemonsets:   make(map[string]*DaemonsetData),
				Statefulsets: make(map[string]*StatefulsetData),
			},
		},
	}

	lg := row.Log()

	// If the status of the apply is deleted, there is case that the encodedID will be exists in the
	// Memory (user delete the application when deployment is running)
	// To create a new db record with deleted status, we add a deleted prefix the the encodedID
	if status == common.ApplyStatusDeleted {
		log.WithField("encoded_id", encodedID).Info("Change encodeding id of the apply")
		encodedID = fmt.Sprintf("deleted-%s", encodedID)
	}
	dr.registryData[encodedID] = &row
	switch status {
	case common.ApplyStatusRunning:
		dr.reporter.DeploymentStarted <- common.DeploymentReport{
			To:          reportTo,
			DeployBy:    deployBy,
			Name:        appName,
			URI:         row.GetURI(),
			Status:      status,
			LogEntry:    lg,
			ClusterName: dr.clusterName,
		}
	case common.ApplyStatusDeleted:
		dr.reporter.DeploymentDeleted <- common.DeploymentReport{
			To:          reportTo,
			DeployBy:    deployBy,
			Name:        appName,
			URI:         row.GetURI(),
			Status:      status,
			LogEntry:    lg,
			ClusterName: dr.clusterName,
		}
	default:
		lg.WithField("status", status).Info("reporter status not supported")
	}

	lg.Info("new application created in registry")

	go row.isFinish(dr.checkFinishDelay)
	return &row

}

// Get will return deployment row that exists in memory
func (dr *RegistryManager) Get(name, namespace, prefix string) *RegistryRow {

	encodedID := generateID(name, namespace, dr.clusterName)
	if prefix != "" {
		encodedID = fmt.Sprintf("%s-%s", prefix, encodedID)

	}
	if row, found := dr.registryData[encodedID]; found {
		return row
	}
	return nil

}

// Log returns the main log entry
func (wbr *RegistryRow) Log() log.Entry {

	lg := log.WithFields(log.Fields{
		"application": wbr.DBSchema.Application,
		"namespace":   wbr.DBSchema.Namespace,
		"cluster":     wbr.DBSchema.Cluster,
		"apply_id":    wbr.GetApplyID(),
	})

	return *lg
}

// GetApplyID generate a uniqe for a specific apply
func (wbr *RegistryRow) GetApplyID() string {

	encodedID := generateID(wbr.DBSchema.Application, wbr.DBSchema.Namespace, wbr.DBSchema.Cluster)
	h := sha1.New()

	h.Write([]byte(fmt.Sprintf("%s-%d", encodedID, wbr.DBSchema.CreationTimestamp)))
	return fmt.Sprintf("%x", h.Sum(nil))

}

// GetURI will generate uri link for UI
func (wbr *RegistryRow) GetURI() string {
	return fmt.Sprintf("deployments/%s/%d", wbr.DBSchema.Application, wbr.DBSchema.CreationTimestamp)

}

// isDeploymentFinish will check for Deployment resource and see if it finished or errord due to timeout.
func (wbr *RegistryRow) isDeploymentFinish() (bool, error) {
	lg := wbr.Log()
	isFinished := false
	diff := time.Now().Sub(time.Unix(wbr.DBSchema.CreationTimestamp, 0)).Seconds()
	if len(wbr.DBSchema.Resources.Deployments) == 0 {
		isFinished = true
		return isFinished, nil
	}
	countOfRunningReplicas := 0
	var desiredStateCount int32
	var readyReplicasCount int32
	for _, deployment := range wbr.DBSchema.Resources.Deployments {
		desiredStateCount = desiredStateCount + deployment.Deployment.DesiredState
		for _, replica := range deployment.Replicaset {
			if replica.Status.Replicas > 0 {
				countOfRunningReplicas = countOfRunningReplicas + 1
			}
			readyReplicasCount = readyReplicasCount + replica.Status.ReadyReplicas
		}
		if deployment.ProgressDeadlineSeconds < int64(diff) {
			lg.WithFields(log.Fields{
				"progress_deadline_seconds": deployment.ProgressDeadlineSeconds,
				"deploy_time":               diff,
			}).Error("deployment failed due to progress deadline")
			return isFinished, errors.New("ProgrogressDeadline has passed")
		}

	}
	lg.WithFields(log.Fields{
		"replicaset_count":     countOfRunningReplicas,
		"desired_state_count":  desiredStateCount,
		"ready_replicas_count": readyReplicasCount,
		"count_deployments":    len(wbr.DBSchema.Resources.Deployments),
	}).Info("deployment status")
	deploymentsNum := len(wbr.DBSchema.Resources.Deployments)
	if deploymentsNum == countOfRunningReplicas && desiredStateCount == readyReplicasCount || wbr.status == common.ApplyStatusDeleted {
		lg.WithFields(log.Fields{
			"replicaset_count":     countOfRunningReplicas,
			"desired_state_count":  desiredStateCount,
			"ready_replicas_count": readyReplicasCount,
		}).Info("deployment has finished successfully")

		// Wating few minutes to collect more event after deployment finished
		isFinished = true
		return isFinished, nil
	}
	return isFinished, nil
}

//isDaemonSetFinish  a DaemonSet is finished if: DesiredNumberScheduled == CurrentNumberScheduled AND DesiredNumberScheduled == UpdatedNumberScheduled
func (wbr *RegistryRow) isDaemonSetFinish() (bool, error) {
	lg := wbr.Log()
	isFinished := false
	if len(wbr.DBSchema.Resources.Daemonsets) == 0 {
		isFinished = true
		return isFinished, nil
	}
	totalDesiredPods := int32(0)
	totalUpdatedPodsOnNodes := int32(0)
	totalCurrentPods := int32(0)
	diff := time.Now().Sub(time.Unix(wbr.DBSchema.CreationTimestamp, 0)).Seconds()
	for _, daemonset := range wbr.DBSchema.Resources.Daemonsets {
		totalDesiredPods = totalDesiredPods + daemonset.Status.DesiredNumberScheduled
		totalUpdatedPodsOnNodes = totalUpdatedPodsOnNodes + daemonset.Status.DesiredNumberScheduled
		totalCurrentPods = totalCurrentPods + daemonset.Status.CurrentNumberScheduled

		if daemonset.ProgressDeadlineSeconds < int64(diff) {
			lg.WithFields(log.Fields{
				"progress_deadline_seconds": daemonset.ProgressDeadlineSeconds,
				"deploy_time":               diff,
			}).Error("daemonset failed due to progress deadline")
			return isFinished, errors.New("ProgrogressDeadline has passed")
		}
	}
	lg.WithFields(log.Fields{
		"total_daemonsets_desired_pods": totalDesiredPods,
		"current_pods_count":            totalCurrentPods,
		"total_daemonsets":              len(wbr.DBSchema.Resources.Daemonsets),
	}).Debug("daemonset status")
	if totalDesiredPods == totalCurrentPods && totalDesiredPods == totalUpdatedPodsOnNodes || wbr.status == common.ApplyStatusDeleted {
		lg.WithFields(log.Fields{
			"total_daemonsets_desired_pods": totalDesiredPods,
			"current_pods_count":            totalCurrentPods,
			"total_daemonsets":              len(wbr.DBSchema.Resources.Daemonsets),
		}).Info("daemonset has finished successfully")
		// Wating few minutes to collect more event after deployment finished
		isFinished = true
		return isFinished, nil
	}
	return isFinished, nil
}

// isStatefulSetFinish defines when a deployment of Statefulset id done.
/* In order to finish a successful deployment you will have to have the following terms:
- Total Pods defined in statefulset yaml should be equal to ready pods running.
- Counts of pods which are committed to the state should be equal to running pods running.
*/
func (wbr *RegistryRow) isStatefulSetFinish() (bool, error) {
	lg := wbr.Log()
	isFinished := false
	diff := time.Now().Sub(time.Unix(wbr.DBSchema.CreationTimestamp, 0)).Seconds()
	if len(wbr.DBSchema.Resources.Statefulsets) == 0 {
		isFinished = true
		return isFinished, nil
	}
	var countOfPodsInState int32
	var countOfRunningPods int32
	var totalDesiredPods int32
	var readyPodsCount int32
	for _, statefulset := range wbr.DBSchema.Resources.Statefulsets {
		totalDesiredPods = statefulset.Statefulset.DesiredState
		countOfRunningPods = countOfRunningPods + statefulset.Status.Replicas
		readyPodsCount = readyPodsCount + statefulset.Status.ReadyReplicas
		countOfPodsInState = int32(len(statefulset.Pods))

		if statefulset.ProgressDeadlineSeconds < int64(diff) {
			lg.WithFields(log.Fields{
				"progress_deadline_seconds": statefulset.ProgressDeadlineSeconds,
				"deploy_time":               diff,
			}).Error("statefulset failed due to progress deadline")
			return isFinished, errors.New("ProgressDeadLine has passed")
		}
	}
	lg.WithFields(log.Fields{
		"total_statefulsets_desired_pods":  totalDesiredPods,
		"total_statefulsets_in_state_pods": countOfPodsInState,
		"current_pods_count":               countOfRunningPods,
		"total_statefulsets":               len(wbr.DBSchema.Resources.Statefulsets),
	}).Info("statefulset status")
	if totalDesiredPods == readyPodsCount && countOfPodsInState == countOfRunningPods || wbr.status == common.ApplyStatusDeleted {
		lg.WithFields(log.Fields{
			"total_statefulset_desired_pods":   totalDesiredPods,
			"total_statefulsets_in_state_pods": countOfPodsInState,
			"current_pods_count":               countOfRunningPods,
			"total_statefulsets":               len(wbr.DBSchema.Resources.Statefulsets),
		}).Info("statefulset has finished successfully")
		// Wating few minutes to collect more event after deployment finished
		isFinished = true
		return isFinished, nil
	}
	return isFinished, nil
}

// isFinish will check (by interval number) when the deployment finished by replicaset status
func (wbr *RegistryRow) isFinish(checkFinishDelay time.Duration) {
	lg := wbr.Log()
	lg.WithFields(log.Fields{
		"deployment_count":   len(wbr.DBSchema.Resources.Deployments),
		"daemonsets_count":   len(wbr.DBSchema.Resources.Daemonsets),
		"statefulsets_count": len(wbr.DBSchema.Resources.Statefulsets),
		"applied_by":         wbr.DBSchema.DeployBy,
		"check_delay":        checkFinishDelay,
	}).Debug("starting to watch on registry row to check if all resources status")
	time.Sleep(checkFinishDelay)

	if wbr.status == common.ApplyStatusDeleted {
		wbr.Stop(common.ApplyStatusDeleted, common.ApplyStatusDescriptionSuccessful)
		return
	}
	for {
		select {
		case <-time.After(time.Second * 2):
			if wbr.finish {
				return
			}
			isDepFinished, depErr := wbr.isDeploymentFinish()
			isDsFinished, dsErr := wbr.isDaemonSetFinish()
			isSsFinished, ssErr := wbr.isStatefulSetFinish()
			if dsErr != nil || depErr != nil || ssErr != nil {
				wbr.Stop(common.ApplyStatusFailed, common.ApplyStatusDescriptionProgressDeadline)
				lg.WithFields(log.Fields{
					"deployment_error":  depErr,
					"daemonset_error":   dsErr,
					"statefulset_error": ssErr,
				}).Error("isFinish function watcher had an error")
				return
			} else if isDepFinished && isDsFinished && isSsFinished {
				wbr.Stop(common.ApplySuccessful, common.ApplyStatusDescriptionSuccessful)
			}
		case <-wbr.ctx.Done():
			lg.Debug("isFinish function watch was stopped, got ctx done signal")
			return

		}
	}
}

// Stop will marked the row as finish
func (wbr *RegistryRow) Stop(status common.DeploymentStatus, message common.DeploymentStatusDescription) {
	lg := wbr.Log()
	lg.WithField("status", status).Debug("marked as done")

	wbr.beforeFinish = true
	time.Sleep(wbr.collectDataAfterDeploymentFinish)
	wbr.DBSchema.DeploymentDescription = message
	wbr.finish = true
	wbr.status = status
	wbr.cancelFn()
}

// UpdateDeploymentStatus will update deployment status
func (dd *DeploymentData) UpdateDeploymentStatus(status appsV1.DeploymentStatus) {
	dd.Status = status
}

// UpdateDeploymentEvents will append events to deployment
func (dd *DeploymentData) UpdateDeploymentEvents(event EventMessages) {
	dd.Events = append(dd.Events, event)
}

// InitReplicaset create new list of replicaset
func (dd *DeploymentData) InitReplicaset(name string) {
	if _, found := dd.Replicaset[name]; !found {
		dd.Replicaset[name] = Replicaset{
			Events: &[]EventMessages{},
			Status: &appsV1.ReplicaSetStatus{},
		}
	}
}

// UpdateReplicasetEvents will append event to replicaset
func (dd *DeploymentData) UpdateReplicasetEvents(name string, event EventMessages) error {
	if _, found := dd.Replicaset[name]; !found {
		return errors.New("Replicaset not found")
	}
	*dd.Replicaset[name].Events = append(*dd.Replicaset[name].Events, event)

	return nil
}

// UpdateReplicasetStatus will update replicaset status
func (dd *DeploymentData) UpdateReplicasetStatus(name string, status appsV1.ReplicaSetStatus) error {
	if _, found := dd.Replicaset[name]; !found {
		return errors.New("Replicaset not found")
	}
	*dd.Replicaset[name].Status = status
	return nil
}

// NewPodToPods
func NewPodToPods(pods map[string]DeploymenPod, pod *v1.Pod) error {
	if _, found := pods[pod.GetName()]; found {

		return errors.New("Pod already exists in pod list")
	}
	phase := string(pod.Status.Phase)
	pods[pod.GetName()] = DeploymenPod{
		Phase:  &phase,
		Events: &[]EventMessages{},
		Pvcs:   map[string][]EventMessages{},
	}
	return nil
}

// NewPod will set new pod to deployment row
func (dd *DeploymentData) NewPod(pod *v1.Pod) error {
	return NewPodToPods(dd.Pods, pod)
}

// UpdatePodEvents will add event to pod events list
func UpdatePodEvents(pods map[string]DeploymenPod, podName string, pvcName string, event EventMessages) error {
	if _, found := pods[podName]; !found {
		log.WithField("pod", podName).Warn("pod does not exist in pod list")
		return errors.New("pod does not exist in pod list")
	}
	// Validate that we not inset duplicated events
	for _, saveEvent := range *pods[podName].Events {
		if saveEvent.Message == event.Message && saveEvent.Time == event.Time {
			return nil
		}
	}

	if pvcName != "" {
		if _, found := pods[podName].Pvcs[pvcName]; !found {
			pods[podName].Pvcs[pvcName] = []EventMessages{}
		}
		pods[podName].Pvcs[pvcName] = append(pods[podName].Pvcs[pvcName], event)
	} else {
		*pods[podName].Events = append(*pods[podName].Events, event)
	}

	return nil
}

// UpdatePodEvents will set pod events
func (dd *DeploymentData) UpdatePodEvents(podName string, pvcName string, event EventMessages) error {
	return UpdatePodEvents(dd.Pods, podName, pvcName, event)
}

// Get the deployment name
func (dd *DeploymentData) GetName() string {
	return dd.Deployment.Name
}

// UpdatePodStatus will change pod status
func UpdatePodStatus(pods map[string]DeploymenPod, pod *v1.Pod, status string) error {
	if _, found := pods[pod.GetName()]; !found {
		log.WithField("pod", pod.GetName()).Warn("pod does not exist in pod list")
		return errors.New("pod does not exist in pod list")
	}
	*pods[pod.GetName()].Phase = status
	return nil
}

// UpdatePod will set pod events to deployment
func (dd *DeploymentData) UpdatePod(pod *v1.Pod, status string) error {
	return UpdatePodStatus(dd.Pods, pod, status)

}

// UpdateApplyStatus will uppdate a daemonsets status
func (dsd *DaemonsetData) UpdateApplyStatus(status appsV1.DaemonSetStatus) {
	dsd.Status = status
}

// UpdateDaemonsetEvents will add event to a daemonset
func (dsd *DaemonsetData) UpdateDaemonsetEvents(event EventMessages) {
	dsd.Events = append(dsd.Events, event)
}

// UpdatePodEvents will set pod events
func (dsd *DaemonsetData) UpdatePodEvents(podName string, pvcName string, event EventMessages) error {
	return UpdatePodEvents(dsd.Pods, podName, pvcName, event)
}

// UpdatePod will set pod events to daemonset
func (dsd *DaemonsetData) UpdatePod(pod *v1.Pod, status string) error {
	return UpdatePodStatus(dsd.Pods, pod, status)
}

// attach a new pod to the daemonset row
func (dsd *DaemonsetData) NewPod(pod *v1.Pod) error {
	return NewPodToPods(dsd.Pods, pod)
}

// GetName will get the daemonset name
func (dsd *DaemonsetData) GetName() string {
	return dsd.Metadata.Name
}

// UpdateStatefulsetEvents will append events to StatefulsetEvents list
func (ssd *StatefulsetData) UpdateStatefulsetEvents(event EventMessages) {
	ssd.Events = append(ssd.Events, event)
}

// UpdatePod will set pod events to statefulset
func (ssd *StatefulsetData) UpdatePod(pod *v1.Pod, status string) error {
	return UpdatePodStatus(ssd.Pods, pod, status)
}

// UpdatePodEvents will set pod events
func (ssd *StatefulsetData) UpdatePodEvents(podName string, pvcName string, event EventMessages) error {
	return UpdatePodEvents(ssd.Pods, podName, pvcName, event)
}

// GetName get the Statefulset name
func (ssd *StatefulsetData) GetName() string {
	return ssd.Statefulset.Name
}

// NewPod Attach a new pod to the Statefulset row
func (ssd *StatefulsetData) NewPod(pod *v1.Pod) error {
	return NewPodToPods(ssd.Pods, pod)
}

// UpdateApplyStatus will update a statefulset status
func (ssd *StatefulsetData) UpdateApplyStatus(status appsV1.StatefulSetStatus) {
	ssd.Status = status
}

// save will save all the row list to the storage
func (dr *RegistryManager) save() {

	dr.saveLock.Lock()
	defer dr.saveLock.Unlock()

	var wg sync.WaitGroup
	wg.Add(len(dr.registryData))
	deleteRows := []string{}
	for key, data := range dr.registryData {
		go func(key string, data *RegistryRow, deleteRows *[]string) {
			defer wg.Done()
			if data.applyID == "" {

				applyID, err := dr.storage.CreateApply(data, data.status)
				if err != nil {
					*deleteRows = append(*deleteRows, key)
					return
				}
				data.applyID = applyID
			} else {
				dr.storage.UpdateApply(data.applyID, data, data.status)
			}

			log.WithFields(log.Fields{
				"name": data.DBSchema.Application,
			}).Debug("deployment was saved")

			if data.finish {

				if data.status != common.ApplyStatusDeleted {
					dr.reporter.DeploymentFinished <- common.DeploymentReport{
						To:          data.DBSchema.ReportTo,
						DeployBy:    data.DBSchema.DeployBy,
						Name:        data.DBSchema.Application,
						URI:         data.GetURI(),
						Status:      data.status,
						LogEntry:    data.Log(),
						ClusterName: dr.clusterName,
					}
				}

				*deleteRows = append(*deleteRows, key)
			}

		}(key, data, &deleteRows)

	}

	wg.Wait()

	for _, key := range deleteRows {
		delete(dr.registryData, key)
	}

}

// generateID will create a id for the deployment
func generateID(name, namespace, cluster string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s-%s-%s", name, namespace, cluster)))
}

// updateAppliesVersionHistory updates a new version of hash kind
func (dr *RegistryManager) updateAppliesVersionHistory(name, namespace, resourceName string, hash uint64) bool {
	return dr.storage.UpdateAppliesVersionHistory(fmt.Sprintf(applyVersionFormat, resourceName, namespace, name, dr.clusterName), hash)
}

// deleteAppliedVersion delete apply version
func (dr *RegistryManager) deleteAppliedVersion(name, namespace, resourceName string) bool {
	return dr.storage.DeleteAppliedVersion(fmt.Sprintf(applyVersionFormat, resourceName, namespace, name, dr.clusterName))
}
