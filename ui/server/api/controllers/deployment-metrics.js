const axios = require('axios');
const config = require('../../config');
const urlPath = `/application/metric`;
const querystring = require('querystring');
const {error} = require('../../logger');
module.exports = {
  urlPath,
  async getAll(params = '') {
    return axios.get(`${config.apiUrl}${urlPath}${params}`)
  },
  async getMetric(metric, provider, deploymentTime, minutesBefore, minutesAfter) {
    const params = {
      provider,
      query: `${metric}`,
      from: (deploymentTime) - (minutesBefore * 60),
      to: (deploymentTime) + (minutesAfter * 60),
    };
    return new Promise(async (resolve, reject) => {
      try {
        const {data} = await axios.get(`${config.apiUrl}${urlPath}?${querystring.stringify(params)}`);
        resolve(data)
      }
      catch (e) {
        error(`cannot get metrics for metric=${metric}, provider=${provider}, deploymentTime=${deploymentTime}`);
        error(e);
        reject(e);
      }
    })
  }
};
