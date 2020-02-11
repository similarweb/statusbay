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
                           cluster = false,
                           from = false,
                           to = false,
                           name = false,
                           nameSpace = false,
                           status = false,
                           userName = false,
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
        namespace: nameSpace.join(','),
        status: status.join(','),
        userName,
        distinct
    };
    return querystring.stringify(Object.fromEntries(Object.entries(query).filter(x => x[1] !== false && x[1] !== '')));
};
module.exports = {
    urlPath,
    async getAll(params) {
        return axios.get(`${config.apiBaseUrl}${urlPath}?${prepareParams(params)}`)
    },
    async getDistinct(params) {
        return axios.get(`${config.apiBaseUrl}${urlPath}?${prepareParams(params)}&distinct=true`)
    }
};
