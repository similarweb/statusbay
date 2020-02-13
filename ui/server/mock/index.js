 const applications = require('../api/controllers/applications');
const clusters = require('../api/controllers/clusters');
const nameSpaces = require('../api/controllers/name-spaces');
const statuses = require('../api/controllers/statuses');
const deploymentDetails = require('../api/controllers/deployment-details');
const config = require('../config');
const fetcher = require('../services/fetcher');
const MockAdapter = require('axios-mock-adapter');
const mockApps = require('./applications.mock');
const mockDeployments = require('./deployments.mock');
const mockClusters = require('./clusters.mock');
const mockNameSpaces = require('./name-spaces.mock');
const mockStatuses = require('./statuses.mock');
const mockDeploymentDetails = require('./deployment-details.mock');
const {getQuery} = require('../services/helpers');

const mock = new MockAdapter(fetcher);
// TODO: add 'distinct' filter
const filter = (params, data) => {
    const {offset, limit, cluster, from, to, name, nameSpace, status, userName} = params;
    let filtered = [...data];
    filtered.sort((a, b) => {
        const {sortDirection, sortBy} = params;
        if (a[sortBy] > b[sortBy]) {
            return sortDirection === 'desc' ? -1 : 1;
        }
        if (b[sortBy] > a[sortBy]) {
            return sortDirection === 'desc' ? 1 : -1;
        }
        return 0;
    });
    if(cluster) filtered = filtered.filter(item => cluster.split(',').includes(item.Cluster));
    if(nameSpace) filtered = filtered.filter(item => nameSpace.split(',').includes(item.Namespace));
    if(status) filtered = filtered.filter(item => status.split(',').includes(item.Status));
    if(userName) filtered = filtered.filter(item => item.DeployBy.includes(userName));
    if(name) filtered = filtered.filter(item => item.Name.includes(name));
    if(from && to) filtered = filtered.filter(item => item.Time <= parseInt(from) && item.Time >= parseInt(to) );

    const limitNum = parseInt(limit);
    const offsetNum = parseInt(offset);

    return {results: filtered.slice(limitNum * offsetNum, limitNum * offsetNum + limitNum), totalCount: filtered.length};
};

mock.onGet(`${config.apiBaseUrl}${clusters.urlPath}`)
    .reply(200, mockClusters.getAll());
mock.onGet(`${config.apiBaseUrl}${nameSpaces.urlPath}`)
    .reply(200, mockNameSpaces.getAll());
 mock.onGet(`${config.apiBaseUrl}${statuses.urlPath}`)
 .reply(200, mockStatuses.getAll());


 mock.onGet(new RegExp(`${config.apiBaseUrl}${applications.urlPath}/*`))
 .reply((config) => {
     const params = getQuery(config.url);
     return [200, filter(params, mockApps.getAll())]
 });

// mock socket data
mock.onGet(`${config.apiBaseUrl}${deploymentDetails.urlPath}`)
.reply(() => {
    return Promise.resolve([200, mockDeploymentDetails.getAll()])
});
