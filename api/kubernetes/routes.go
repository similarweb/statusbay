package kubernetes

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"statusbay/api/httpresponse"
	"statusbay/config"
	"statusbay/state"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

type RouterKubernetesManager struct {
	storage          Storage
	router           *mux.Router
	eventMarksConfig config.KubernetesMarksEvents
}

//NewKubernetesRoutes sets up the Kubernetes router to handle API endpoints
func NewKubernetesRoutes(storage Storage, router *mux.Router, eventPath string) *RouterKubernetesManager {

	eventMarksConfig, err := config.LoadKubernetesMarksConfig(eventPath)
	if err != nil {
		log.WithError(err).WithField("path", eventPath).Error("could not load events configuration file")
	}

	kubernetesRoutes := &RouterKubernetesManager{
		storage:          storage,
		router:           router,
		eventMarksConfig: eventMarksConfig,
	}
	kubernetesRoutes.bindEndpoints()
	return kubernetesRoutes
}

//bindEndpoints List of API endpoints
func (kr *RouterKubernetesManager) bindEndpoints() {
	kr.router.HandleFunc("/api/v1/kubernetes/applications", kr.Applications).Methods("GET")
	kr.router.HandleFunc("/api/v1/kubernetes/applications/values/{column}", kr.ApplicationsColumnValues).Methods("GET")
	kr.router.HandleFunc("/api/v1/kubernetes/application/{apply_id}", kr.GetDeployment).Methods("GET")
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
		httpresponse.JSONWrite(resp, http.StatusNotFound, httpresponse.HTTPErrorResponse{Error: "Could not return applications"})
	}

	count, err := route.storage.ApplicationsCount(allFilter)
	if err != nil {
		httpresponse.JSONWrite(resp, http.StatusNotFound, httpresponse.HTTPErrorResponse{Error: "Could not return all applications"})
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
		errs.Add("column", "Column name is not allow")
	}

	if len(errs) > 0 {
		httpresponse.JSONErrorParameters(resp, http.StatusBadRequest, errs)
		return
	}

	var table *state.TableKubernetes

	values, err := route.storage.GetUniqueFieldValues(table.TableName(), columnName)
	if err != nil {
		httpresponse.JSONError(resp, http.StatusNotFound, err)

	}

	httpresponse.JSONWrite(resp, http.StatusOK, values)
}

//GetDeployment returns a specific deployment details.
func (route *RouterKubernetesManager) GetDeployment(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	applyID := params["apply_id"]

	deployment, err := route.storage.GetDeployment(applyID)
	if err != nil {
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Deployment not found"))
		return
	}

	var kubernetesDeploymentResponse ResponseDeploymentData

	err = json.Unmarshal([]byte(deployment.Details), &kubernetesDeploymentResponse)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{}).Error("Could not parse deployment detail")
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Could not parse deployment detail"))
		return
	}

	MarkApplicationDeploymentEvents(&kubernetesDeploymentResponse, route.eventMarksConfig)

	response := ResponseKubernetesDeployment{
		Name:      deployment.Name,
		Time:      deployment.Time,
		Status:    deployment.Status,
		Cluster:   deployment.Cluster,
		Namespace: deployment.Namespace,
		Details:   kubernetesDeploymentResponse,
	}

	httpresponse.JSONWrite(resp, http.StatusOK, response)

}
