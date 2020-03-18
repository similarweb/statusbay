const axios = require('axios');
const config = require('../../config');
const urlPath = `/applications`;
const querystring = require('querystring');

axios.interceptors.request.use(request => {
  return request
})

const prepareParams = ({
                         page = 0,
                         rowsPerPage = 10,
                         sortBy = "Name",
                         sortDirection = "desc",
                         cluster = [],
                         from = false,
                         to = false,
                         name = false,
                         exactName = false,
                         nameSpace = [],
                         status = [],
                         deployBy = false,
                         distinct = false
                       } = {}) => {
  const query = {
    offset: page,
    limit: rowsPerPage,
    sortby: sortBy,
    sortdirection: sortDirection,
    cluster: cluster.join(','),
    from,
    to,
    name,
    exactName,
    namespace: nameSpace.join(','),
    status: status.join(','),
    deployby: deployBy,
    distinct
  };
  return querystring.stringify(Object.fromEntries(Object.entries(query)
  .filter(x => x[1] !== false && x[1] !== '')));
};
module.exports = {
  urlPath,
  async getAll(params) {
    return axios.get(`${config.kubernetesApiUrl}${urlPath}?${prepareParams(params)}`)
  },
  async getDistinct(params) {
    return axios.get(`${config.kubernetesApiUrl}${urlPath}?${prepareParams(params)}&distinct=true`)
  }
};
