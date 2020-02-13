const axios = require('axios');
const config = require('../../config');
const urlPath = `/applications/values/status`;
module.exports = {
  urlPath,
  async getAll() {
    return axios.get(`${config.apiBaseUrl}${urlPath}`)
  }
};
