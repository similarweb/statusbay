package kubernetes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"statusbay/api/httpresponse"
	"statusbay/config"
	"statusbay/state"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type RouterKubernetesManager struct {
	storage             Storage
	router              *mux.Router
	eventMarksConfig    config.KubernetesMarksEvents
	absoluteLogsPodPath string
}

//NewKubernetesRoutes sets up the Kubernetes router to handle API endpoints
func NewKubernetesRoutes(storage Storage, router *mux.Router, eventsConfig config.KubernetesMarksEvents, absoluteLogsPodPath string) *RouterKubernetesManager {
	kubernetesRoutes := &RouterKubernetesManager{
		storage:             storage,
		router:              router,
		eventMarksConfig:    eventsConfig,
		absoluteLogsPodPath: absoluteLogsPodPath,
	}
	kubernetesRoutes.bindEndpoints()
	return kubernetesRoutes
}

//bindEndpoints List of API endpoints
func (kr *RouterKubernetesManager) bindEndpoints() {
	kr.router.HandleFunc("/api/v1/kubernetes/applications", kr.Applications).Methods("GET")
	kr.router.HandleFunc("/api/v1/kubernetes/applications/values/{column}", kr.ApplicationsColumnValues).Methods("GET")
	kr.router.HandleFunc("/api/v1/kubernetes/application/{apply_id}", kr.GetDeployment).Methods("GET")
	kr.router.HandleFunc("/api/v1/kubernetes/application/{apply_id}/logs/pod/{pod}", kr.PodLogs).Methods("GET")
}

//Applications returns a list of applied application.
func (route *RouterKubernetesManager) Applications(resp http.ResponseWriter, req *http.Request) {

	// Applications' filter
	queryFilter := FilterApplication(req)

	// Number of results filter (excluding limit and offset)
	unlimitedReq := req
	unlimitedReqQuery := unlimitedReq.URL.Query()
	unlimitedReqQuery.Set("limit", "-1")
	unlimitedReqQuery.Set("offset", "0")
	unlimitedReq.URL.RawQuery = unlimitedReqQuery.Encode()

	allFilter := FilterApplication(unlimitedReq)

	rows, err := route.storage.Applications(queryFilter)
	if err != nil {
		log.WithError(err).Error("could not return filtered applications")
		httpresponse.JSONWrite(resp, http.StatusNotFound, httpresponse.HTTPErrorResponse{Error: "Could not return applications"})
		return
	}

	count, err := route.storage.ApplicationsCount(allFilter)
	if err != nil {
		log.WithError(err).Error("could not return all applications")
		httpresponse.JSONWrite(resp, http.StatusNotFound, httpresponse.HTTPErrorResponse{Error: "Could not return all applications"})
		return
	}

	response := []ResponseKubernetesApplications{}

	for _, row := range *rows {

		response = append(response, ResponseKubernetesApplications{
			ApplyID:   row.ApplyId,
			Name:      row.Name,
			Cluster:   row.Cluster,
			Namespace: row.Namespace,
			DeployBy:  row.DeployBy,
			Status:    row.Status,
			Time:      row.Time,
		})

	}

	r := &ResponseKubernetesApplicationsCount{
		Records: response,
		Count:   count,
	}

	httpresponse.JSONWrite(resp, http.StatusOK, r)
}

//ApplicationsColumnValues returns a unique column values
func (route *RouterKubernetesManager) ApplicationsColumnValues(resp http.ResponseWriter, req *http.Request) {

	errs := url.Values{}
	params := mux.Vars(req)
	columnName := params["column"]

	allowColumns := map[string]struct{}{
		"name":      {},
		"cluster":   {},
		"namespace": {},
		"status":    {},
		"deploy_by": {},
	}

	if _, ok := allowColumns[columnName]; !ok {
		errs.Add("column", "Unknown column name")
	}

	if len(errs) > 0 {
		httpresponse.JSONErrorParameters(resp, http.StatusBadRequest, errs)
		return
	}

	var table *state.TableKubernetes

	values, err := route.storage.GetUniqueFieldValues(table.TableName(), columnName)
	if err != nil {
		log.WithFields(log.Fields{
			"column_name": columnName,
			"table_name":  table.TableName(),
		}).Error("could not find column in table")
		httpresponse.JSONError(resp, http.StatusNotFound, err)
		return
	}

	httpresponse.JSONWrite(resp, http.StatusOK, values)
}

//GetDeployment returns a specific deployment details.
func (route *RouterKubernetesManager) GetDeployment(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	applyID := params["apply_id"]

	deployment, err := route.storage.GetDeployment(applyID)
	if err != nil {
		log.WithField("apply_id", applyID).Error("deployment not found")
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Deployment not found"))
		return
	}

	var kubernetesDeploymentResponse ResponseDeploymentData

	err = json.Unmarshal([]byte(deployment.Details), &kubernetesDeploymentResponse)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{}).Error("could not parse deployment details")
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Could not parse deployment detail"))
		return
	}

	MarkApplicationDeploymentEvents(&kubernetesDeploymentResponse, route.eventMarksConfig)

	response := ResponseKubernetesDeployment{
		Name:      deployment.Name,
		DeployBy:  deployment.DeployBy,
		Time:      deployment.Time,
		Status:    deployment.Status,
		Cluster:   deployment.Cluster,
		Namespace: deployment.Namespace,
		Details:   kubernetesDeploymentResponse,
	}

	httpresponse.JSONWrite(resp, http.StatusOK, response)

}

//PodLogs return pod continaers logs
func (route *RouterKubernetesManager) PodLogs(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	applyID := params["apply_id"]
	podName := params["pod"]

	logsFolder := fmt.Sprintf("%s/%s/%s", route.absoluteLogsPodPath, applyID, podName)

	files, err := ioutil.ReadDir(logsFolder)
	if err != nil {
		log.WithField("file", logsFolder).Warn("log file not found")
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("logs not found"))
		return
	}

	responseContainerLogs := []ResponseContainerLogs{}
	for _, file := range files {
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", logsFolder, file.Name()))
		if err != nil {
			log.WithFields(log.Fields{
				"logs_folder": logsFolder,
				"log_file":    file,
			}).Warn("log file not found")
			continue
		}

		responseContainerLogs = append(responseContainerLogs, ResponseContainerLogs{
			ContainerName: file.Name(),
			Lines:         strings.Split(string(data), "\n"),
		})
	}

	httpresponse.JSONWrite(resp, http.StatusOK, responseContainerLogs)

}
