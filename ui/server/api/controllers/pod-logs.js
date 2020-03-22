const axios = require('axios');
const config = require('../../config');
module.exports = {
    async getAll(deploymentId, podName) {
        return axios.get(`${config.kubernetesApiUrl}/application/${deploymentId}/logs/pod/${podName}`)
    }
};
