const axios = require('axios');
const config = require('../../config');
const urlPath = `/latest-version/statusbay`;
module.exports = {
    urlPath,
    async getAll() {
        return axios.get(`${config.apiUrl}${urlPath}`)
    }
};
