const metricsController = require('../api/controllers/deployment-metrics');
const { info, error } = require('../logger');
const metricsTransformer = require('../services/data-transformers/metrics');

const init = (io) => {
  const metrics = io.of('/metrics');
  const emitOnce = async (socket, { metric, provider, deploymentTime }) => {
    info('sending metrics data...');
    try {
      const data = await metricsController.getSingleMetric(metric, provider, deploymentTime);
      const tranformedData = metricsTransformer(data);
      socket.emit('data', {
        data: tranformedData, config: {
          query: metric,
          provider
        }
      });
    }
    catch (e) {
      error(e);
      error(`error getting metrics for ${metric} ${provider} ${deploymentTime}`);
      socket.emit('metric-error', {
        error: {
          code: e.response.code,
          message: e.response.data,
          name: e.name,
          url: e.config.url
        }, config: {
          query: metric,
          provider
        }
      });
    }
  };
  metrics.on('connection', (socket) => {
    let intervalId;
    info('User connected to metrics NS');
    socket.on('init', async (metric, provider, deploymentTime) => {
      info(`metrics init: ${metric} ${provider} ${deploymentTime}`);
      emitOnce(socket, { metric, provider, deploymentTime });
      intervalId = setInterval(async () => {
        emitOnce(socket, { metric, provider, deploymentTime });
      }, 5000)
    });
    socket.on('disconnect', () => {
      clearInterval(intervalId);
      info('disconnect from metric NS');
    })
    socket.on('close', () => {
      clearInterval(intervalId);
      info('close metric NS');
    })
  })
};
module.exports = init;
