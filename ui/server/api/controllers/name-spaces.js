const axios = require('axios');
const config = require('../../config');
const urlPath = `/applications/values/namespace`;
module.exports = {
    urlPath,
    async getAll() {
        return axios.get(`${config.kubernetesApiUrl}${urlPath}`)
    }
};
