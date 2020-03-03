const winston = require('winston');
const GelfTransport = require('winston-gelf');
const config = require('../config')

const [loggerUrl,loggerPort] = config.gelfUrl.split(':');

const logger = winston.createLogger({
  level: config.logLevel,
  format: winston.format.combine(
    winston.format.splat(),
    winston.format.errors({ stack: true }),
    winston.format.simple()
  ),
  transports: config.nodeEnv === 'development' ?
    [
      new winston.transports.Console()
    ]
    : [
      new winston.transports.Console(),
      new GelfTransport({
        gelfPro: {
          fields: {
            application: 'statusbay-ui',
            environment: config.nodeEnv,
          },
          adapterOptions: {
            host: loggerUrl,
            port: loggerPort,
          }
        }
      })
    ]
});

module.exports = logger;
