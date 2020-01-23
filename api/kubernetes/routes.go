package kubernetes

import (
	"encoding/json"
	"errors"
	"net/http"
	"statusbay/config"
	"statusbay/api/httpresponse"

	"strconv"

	"github.com/apex/log"
	"github.com/gorilla/mux"
)

type RouterKubernetesManager struct {
	storage          Storage
	router           *mux.Router
	eventMarksConfig config.KubernetesMarksEvents
}

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

func (kr *RouterKubernetesManager) bindEndpoints() {
	kr.router.HandleFunc("/api/v1/kubernetes/applications", kr.Applications).Methods("GET")
	kr.router.HandleFunc("/api/v1/kubernetes/application/{job_id}/{time}", kr.GetDeployment).Methods("GET")
}

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

func (route *RouterKubernetesManager) GetDeployment(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	applicationName := params["name"]

	deploymentTime, err := strconv.ParseInt(params["time"], 10, 64)
	if err != nil {
		httpresponse.JSONError(resp, http.StatusBadRequest, errors.New("Invalid time parameter"))
		return
	}

	deployment, err := route.storage.GetDeployment(applicationName, deploymentTime)
	if err != nil {
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Deployment not found"))

	}

	var kubernetesDeploymentResponse ResponseDeploymentData

	err = json.Unmarshal([]byte(deployment.Details), &kubernetesDeploymentResponse)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{}).Error("Could not parse deployment detail.")
		httpresponse.JSONError(resp, http.StatusNotFound, errors.New("Could not parse deployment detail."))
	}

	MarkApplicationDeploymentEvents(&kubernetesDeploymentResponse, route.eventMarksConfig)

	response := ResponseKubernetesDeployment{
		Name:      deployment.Name,
		Status:    deployment.Status,
		Cluster:   deployment.Cluster,
		Namespace: deployment.Namespace,
		Details:   kubernetesDeploymentResponse,
	}

	httpresponse.JSONWrite(resp, http.StatusOK, response)

}
