const expect = require('chai').expect;
const nameSpaces = require('../api/controllers/name-spaces');
require('../mock/index');

describe('NameSpaces controller', () => {
    it('gets clusters list and checks object structure', async () => {
        // Arrange

        // Act
        const {data} = await nameSpaces.getAll();

        // Assert
        const {0:first} = data;
        expect(first.hasOwnProperty('Name')).to.be.true;
    });
});
