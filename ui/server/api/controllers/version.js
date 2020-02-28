const axios = require('axios');
const config = require('../../config');
const urlPath = `/version`;
module.exports = {
    urlPath,
    async getAll() {
        return axios.get(`${config.apiUrl}${urlPath}`)
    }
};
