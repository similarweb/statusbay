import { http } from './request.service'

export const jobService = {
    jobs,
    deployments,
    deployment,
    JobNames,
};

/**
 * Getting list of last deployment jobs
 * @param {int} top - number of last deployments
 */
function jobs(top) {
    return http.send(`api/v1/deployments?top=${top}`, `get`).then(this.handleResponse).then(response => {
        return response;
    })
} 

/**
 * Getting list of deployment by job id
 * @param {string} id - nomad job name
 * @param {int} top - number of last deployments
 */
function deployments(id,top) {
    return http.send(`api/v1/deployments/${id}?top=${top}`, `get`).then(this.handleResponse).then(response => {
        return response;
    })
}

/**
 * Getting deployment status
 * @param {string} job - nomad job name
 * @param {id} time - deploy time
 */
function deployment(job, time) {
    return http.send(`api/v1/deployments/${job}/${time}`, `get`).then(this.handleResponse).then(response => {
        return response;
    })
}

function JobNames() {
    return http.send(`api/v1/jobs`, `get`).then(this.handleResponse).then(response => {
        return response;
    })
}
