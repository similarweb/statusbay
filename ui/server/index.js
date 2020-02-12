require('dotenv').config();
const express = require('express');
const app = express();
const api = require('./api/routes');
const port = process.env.SERVER_PORT || 5000;
const socket = require('./socket');
const {info} = require('./logging');
const axios = require('axios');

if (process.env.NODE_ENV === 'development' && !process.env.API_URL) {
    require('./mock');
}

axios.interceptors.request.use(request => {
    info(`Starting Request: ${request.url}`);
    return request
})

app.use('/api', api);

const server = app.listen(port, err => {
    if (err) {
        throw err;
    }
    info(`Server started... Listening at http://localhost:${port}`);
    info(`API_URL: ${process.env.API_URL}`);
    info(`LOGGER_URL: ${process.env.LOGGER_URL}`);
});

socket(server);
