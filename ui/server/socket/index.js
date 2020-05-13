const socketIO = require('socket.io');

const initDeploymentDetails = require('./deployment-details');
const initApplications = require('./applications');
const initMetrics = require('./metrics');
const initAlerts = require('./alerts');
const initPodLogs = require('./pod-logs');

const init = (server) => {
  const io = socketIO(server, {
    path: '/api/socket'
  });
  initDeploymentDetails(io);
  initApplications(io);
  initMetrics(io);
  initPodLogs(io);
}

module.exports = init;
