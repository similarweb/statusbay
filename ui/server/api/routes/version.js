const express = require('express');
const router = express.Router();
const version = require('../controllers/version');
const {error} = require('../../logger');

router.get('/', async (req, res) => {
  try{
    const {data: versionData} = await version.getAll();
    res.status(200).send(versionData);
  }catch (e) {
    error(`cannot get version ${JSON.stringify(e)}`);
    res.status(500).send({error: e.message})
  }
});
module.exports = router;
