const url = require('url');
const querystring = require('querystring');

const getQuery = (fullUrl) => querystring.parse(url.parse(fullUrl).query);

module.exports = {getQuery};
