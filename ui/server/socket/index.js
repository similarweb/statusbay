const socketIO = require('socket.io');

const initDeploymentDetails = require('./deployment-details');
const initApplications = require('./applications');
const initMetrics = require('./metrics');
const initAlerts = require('./alerts');

const init = (server) => {
  const io = socketIO(server, {
    path: '/api/socket'
  });
  initDeploymentDetails(io);
  initApplications(io);
  initMetrics(io);
  initAlerts(io);
}

module.exports = init;
