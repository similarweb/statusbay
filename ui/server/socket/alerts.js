const alertsController = require('../api/controllers/deployment-alerts');
const { info, error } = require('../logger');
const alertsTransformer = require('../services/data-transformers/alerts');
const moment = require('moment');

const createErrorMessage = response => {
  try {
    return Object.values(response.data.Validation).reduce((result, currentValue = []) => {
      result.push(...currentValue);
      return result;
    }, []);
  } catch (e) {
    return ''
  }
}

const init = (io) => {
  const alerts = io.of('/alerts');
  alerts.on('connection', (socket) => {
    let intervalId;
    const emitOnce = async (socket, { tags, provider, deploymentTime }) => {
      info('sending alerts data...');
      try {
        const data = await alertsController.getAll(tags, provider, deploymentTime);
        // empty array means: check tags not found
        if (Array.isArray(data) && data.length === 0) {
          socket.emit('alerts-tags', {
            tags,
            provider
          });
        }
        const transformedData = alertsTransformer(data);
        socket.emit('data', {
          data: transformedData,
          config: {
            tags,
            provider
          }
        });
        // stop calling the API if we got enough data
        const now = moment();
        if (now.isAfter(moment.unix(deploymentTime).add(30, 'minutes'))) {
          info(`stop calling the API for alerts, provider: ${provider}, tags: ${tags}, now: ${now.format('DD/MM/YYYY HH:mm:ss')}, deployment time: ${deploymentTime}`);
          clearInterval(intervalId);
        }
      }
      catch (e) {
        error(e);
        error(`error getting alerts for ${tags} ${provider} ${deploymentTime}`);
        if (e.response.status === 400) {
          info('stop calling alerts endpoint: status code: 400');
          socket.emit('alerts-error', {
            error: {
              code: e.response.status,
              message: createErrorMessage(e.response),
              url: e.config.url
            }
          })
          clearInterval(intervalId);
        }
      }
    };
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
