const axios = require('axios');
const config = require('../../config');
const urlPath = `/application/`;
module.exports = {
    urlPath,
    async getAll(id) {
        return axios.get(`${config.kubernetesApiUrl}${urlPath}${id}`)
    }
};
