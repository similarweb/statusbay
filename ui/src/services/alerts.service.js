import { http } from './request.service'

export const alertsService = {
    alerts,
};

/**
 * Get alerts
 * @param {string} tags - tags filter 
 * @param {int} from  - unix timestamp
 * @param {int} to - unix timestamp
 */
function alerts(tags, from, to) {

    return http.send(`api/v1/alerts?tags=${tags}&from=${from}&to=${to}`, `get`).then(this.handleResponse).then(response => {
        return response;
    })

}