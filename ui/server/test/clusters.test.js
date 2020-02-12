const expect = require('chai').expect;
const clusters = require('../api/controllers/clusters');
require('../mock/index');

describe('Clusters controller', () => {
    it('gets clusters list and checks object structure', async () => {
        // Arrange

        // Act
        const {data} = await clusters.getAll();

        // Assert
        const {0:first} = data;
        expect(first.hasOwnProperty('Name')).to.be.true;
    });
});
