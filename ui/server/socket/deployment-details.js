const detailsController = require('../api/controllers/deployment-details');
const { info, error } = require('../logger');
const { convertDeploymentDetailsData } = require('../services/data-transformers/deploymet-details');
const hash = require('object-hash');

const init = (io) => {
  const deploymentDetails = io.of('/deployment-details');
  deploymentDetails.on('connection', (socket) => {
    let intervalId;
    info('User connected to deploymentDetails NS');
    const emitOnce = async (socket, id) => {
      info('sending deploymentDetails data...');
      try {
        const { data } = await detailsController.getAll(id);
        const tranformedData = convertDeploymentDetailsData(data);
        const hashValue = hash(tranformedData);
        socket.emit('data', { data: tranformedData, hashValue });
        // stop interval if status isn't running since the data won't change anymore
        // TODO: consider closing the socket
        if (tranformedData && tranformedData.status !== 'running') {
          if (intervalId) {
            info(`stop getting deployment details, status is ${tranformedData.status}`);
            clearInterval(intervalId);
          }
        }
      }
      catch (e) {
        error(`error getting deployments details for ${id} error ${e}`);
        if (e.response && e.response.status === 404) {
          socket.emit('not-found', { error: { code: 404, url: e.request.path } });
          clearInterval(intervalId);
        }
      }
    };
    socket.on('init', async (id) => {
      info(`deploymentDetails init: ${id}`);
      emitOnce(socket, id);
      intervalId = setInterval(async () => {
        emitOnce(socket, id);
      }, 2000)

    });
    socket.on('disconnect', () => {
      clearInterval(intervalId);
      info('disconnect from deployment details NS');
    })
    socket.on('close', () => {
      clearInterval(intervalId);
      info('close deployment details NS');
    })
  })
};
module.exports = init;
