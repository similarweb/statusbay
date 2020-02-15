const {getQuery} = require("../services/helpers");
const expect = require('chai').expect;
const applications = require('../api/controllers/applications');
require('../mock/index');

describe('Gets Applications list and', () => {
    it('checks default params', async () => {
        // Act
        const {data, config} = await applications.getAll();
        const query = getQuery(config.url);

        // Assert
        expect(data.length).to.equal(100);
        expect(query).to.deep.equal({"limit":"10", offset: "0", sortBy: "Name", sortDirection: "desc"})
    });
    it('gets applications list and checks object structure', async () => {
        // Arrange

        // Act
        const {data} = await applications.getAll();

        // Assert
        const {0:first} = data;
        expect(first.hasOwnProperty('Name')).to.be.true;
        expect(first.hasOwnProperty('Status')).to.be.true;
        expect(first.hasOwnProperty('Cluster')).to.be.true;
        expect(first.hasOwnProperty('Namespace')).to.be.true;
        expect(first.hasOwnProperty('DeployBy')).to.be.true;
        expect(first.hasOwnProperty('Time')).to.be.true;
    });
    it('gets paginated results', async () => {
        const {data, config} = await applications.getAll({
            page:1,
            rowsPerPage: 20,
            sortBy: "Status|asc",
            filters:"cluster:serving|namespace:web"
        });
        const query = getQuery(config.url);
        expect(query).to.deep.equal({
            offset: '1',
            limit: '20',
            sortBy: 'Status',
            sortDirection: 'asc',
            cluster: 'serving',
            namespace: 'web' });
    })
});
