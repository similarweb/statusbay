const applicationsController = require('../api/controllers/applications');
const { info, error } = require('../logger');
const hash = require('object-hash');

const init = (io) => {
  const emitOnce = async (socket, filters) => {
    try {
      const {data} = await applicationsController.getAll(filters);
      const hashValue = hash(data);
      socket.emit('data', { data, filters, hashValue });
    } catch (e) {
      error(`error getting applications`, {filters, error: e});
    }
  };

  const applications = io.of('/applications');
  applications.on('connection', (socket) => {
    let intervalId;
    info('User connected to applications NS');
    socket.on('filters', async(newFilters) => {
      info('User change filters: ', newFilters);
      info('Clear interval');
      clearInterval(intervalId);
      emitOnce(socket, newFilters);
      intervalId = setInterval(async () => {
        emitOnce(socket, newFilters);
      }, 2000)
    });
    socket.on('disconnect', () => {
      clearInterval(intervalId);
      info('disconnect from applications NS');
    });
    socket.on('close', () => {
      clearInterval(intervalId);
      info('close applications NS');
    })
  })
}
module.exports = init;
