import { http } from './request.service'

export const metricService = {
    metric,
};

/**
 * Getting metric data from the server
 * 
 * @param {string} query url request
 * @param {int} from from date
 * @param {int} to to date
 * @returns {Promise}
 */
function metric(query, from, to) {

    return http.send(`api/v1/metric?query=${query}&from=${from}&to=${to}`, `get`).then(this.handleResponse).then(response => {
        return response;
    })

}