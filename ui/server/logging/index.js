const { trace, error, fatal, info, warn } = require('noderus');
const logLevel = process.env.LOG_LEVEL || 'info';

module.exports = {
  trace,
  error,
  fatal,
  info,
  warn
}
