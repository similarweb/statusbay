const alertsController = require('../api/controllers/deployment-alerts');
const { info, error } = require('../logger');
const alertsTransformer = require('../services/data-transformers/alerts');

const init = (io) => {
  const alerts = io.of('/alerts');
  const emitOnce = async (socket, { tags, provider, deploymentTime }) => {
    info('sending alerts data...');
    try {
      const data = await alertsController.getAll(tags, provider, deploymentTime);
      const tranformedData = alertsTransformer(data);
      socket.emit('data', {
        data: tranformedData,
        config: {
          tags,
          provider
        }
      });
    }
    catch (e) {
      error(e);
      error(`error getting alerts for ${tags} ${provider} ${deploymentTime}`);
    }
  };
  alerts.on('connection', (socket) => {
    let intervalId;
    info('User connected to alerts NS');
    socket.on('init', async (tags, provider, deploymentTime) => {
      info(`alerts init: ${tags} ${provider} ${deploymentTime}`);
      emitOnce(socket, { tags, provider, deploymentTime });
      intervalId = setInterval(async () => {
        emitOnce(socket, { tags, provider, deploymentTime });
      }, 2000)
    });
    socket.on('disconnect', () => {
      clearInterval(intervalId);
      info('disconnect from alerts NS');
    })
    socket.on('close', () => {
      clearInterval(intervalId);
      info('close alerts NS');
    })
  })
};
module.exports = init;
