const detailsController = require('../api/controllers/deployment-details');
const { info, error } = require('../logger');
const { convertDeploymentDetailsData } = require('../services/data-transformers/deploymet-details');

const init = (io) => {
  const deploymentDetails = io.of('/deployment-details');
  const emitOnce = async (socket, id) => {
    info('sending deploymentDetails data...');
    try {
      const { data } = await detailsController.getAll(id);      
      const tranformedData = convertDeploymentDetailsData(data);
      socket.emit('data', { data: tranformedData });
    }
    catch (e) {
      error(`error getting deployments details for ${id} error ${e}`);
    }
  };
  deploymentDetails.on('connection', (socket) => {
    let intervalId;
    info('User connected to deploymentDetails NS');
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
