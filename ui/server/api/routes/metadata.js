const express = require('express');
const router = express.Router();
const clusters = require('../controllers/clusters');
const nameSpaces = require('../controllers/name-spaces');
const statuses = require('../controllers/statuses');
const {error} = require('../../logger');

router.get('/', async (req, res) => {
    try{
        const {data: allClusters} = await clusters.getAll();
        const {data: allNamespaces} = await nameSpaces.getAll();
        const {data: allStatuses} = await statuses.getAll();
        res.status(200).send({allClusters, allNamespaces, allStatuses});
    }catch (e) {
        error(`cannot get metadata ${JSON.stringify(e)}`);
        res.status(500).send({error: e.message})
    }
});
module.exports = router;

