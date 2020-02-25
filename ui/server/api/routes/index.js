const express = require('express');
const router = express.Router();
const applications = require('./applications');
const deployments = require('./deployments');
const metadata = require('./metadata');
const version = require('./version');

router.get('/health', (req, res) => {
    res.status(200).send({ status: 'ok' });
});

router.use('/applications', applications);
router.use('/deployments', deployments);
router.use('/metadata', metadata);
router.use('/version', version);

module.exports = router;
