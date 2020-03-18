const podLogsController = require('../api/controllers/pod-logs');
const { info, error } = require('../logger');
const podLogsTransformer = require('../services/data-transformers/pod-logs');
const moment = require('moment');


const init = (io) => {
  const podLogs = io.of('/pod-logs');
  podLogs.on('connection', (socket) => {
    let intervalId;
    const emitOnce = async (socket, { deploymentId, podName }) => {
      info('sending pod logs data...');
      try {
        const data = await podLogsController.getAll(deploymentId, podName);
        const transformedData = podLogsTransformer(data);
        socket.emit('data', {
          data: transformedData, config: {
            deploymentId,
            podName
          }
        });
      }
      catch (e) {
        error(e);
        error(`error getting pod logs for ${deploymentId} ${podName}`);
        socket.emit('pod-logs-error', {
          error: {
            code: e.response.code,
            message: e.response.data,
            name: e.name,
            url: e.config.url
          }, config: {
            deploymentId,
            podName
          }
        });
      }
    };
    info('User connected to pod logs NS');
    socket.on('init', async (deploymentId, podName) => {
      info(`pod logs init: ${deploymentId} ${podName}`);
      emitOnce(socket, { deploymentId, podName });
      intervalId = setInterval(async () => {
        emitOnce(socket, { deploymentId, podName });
      }, 2000)
    });
    socket.on('disconnect', () => {
      clearInterval(intervalId);
      info('disconnect from pod logs NS');
    })
    socket.on('close', () => {
      clearInterval(intervalId);
      info('close pod logs NS');
    })
  })
};
module.exports = init;
