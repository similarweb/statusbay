const path = require('path');
const express = require('express');
const compression = require('compression');

module.exports = function addProdMiddlewares(app, options) {

  app.use(compression());
  app.use(options.publicPath, express.static(options.outputPath));
  app.get('*', (req, res) => res.sendFile(path.resolve(options.outputPath, 'index.html')));
};
