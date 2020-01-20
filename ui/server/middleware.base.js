
module.exports = (app, options) => {
  const isProd = process.env.NODE_ENV === 'production';
  if (isProd) {
    const addProdMiddlewares = require('./middleware.production');
    addProdMiddlewares(app, options);
  } else {
    const webpackConfig = require('../config/webpack.config.development.js');
    const middleware = require('./middleware.development.js');
    middleware(app, webpackConfig);
  }
  };
  