const socketIO = require('socket.io');

const initDeploymentDetails = require('./deployment-details');
const initApplications = require('./applications');

const init = (server) => {
  const io = socketIO(server, {
    path: '/api/socket'
  });
  initDeploymentDetails(io);
  initApplications(io);
}

module.exports = init;
