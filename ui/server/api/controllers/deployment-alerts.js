const axios = require('axios');
const config = require('../../config');
const urlPath = `/application/alerts`;
const querystring = require('querystring');
const { error } = require('../../logger');
module.exports = {
  urlPath,
  async getAll(tags, provider, deploymentTime, minutesBefore = 30, minutesAfter = 30) {
    const params = {
      provider,
      tags,
      from: (deploymentTime) - (minutesBefore * 60),
      to: (deploymentTime) + (minutesAfter * 60),
    };
    return new Promise(async (resolve, reject) => {
      try {
        const response = await axios.get(`${config.metricsApiUrl}${urlPath}?${querystring.stringify(params)}`);
        const {data} = response;
        resolve(data)
      }
      catch (e) {
        error(`cannot get alerts for tags=${tags}, provider=${provider}, deploymentTime=${deploymentTime}`);
        error(e);
        resolve([])
      }
    })
  }
};
