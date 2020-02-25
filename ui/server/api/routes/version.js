const express = require('express');
const router = express.Router();
const version = require('../controllers/version');
const {error} = require('../../logger');

router.get('/', async (req, res) => {
  try{
    // const versionData = await version.getAll();
    const versionData = {
      "LatestVersion": "1.0.0",
      "LatestReleaseDate": 12345,
      "Outdated": true,
      "Notifications": [
        {
          "date": 1234,
          "message": "Link to the version...."
        }
      ]
    }
    res.status(200).send(versionData);
  }catch (e) {
    error(`cannot get version ${JSON.stringify(e)}`);
    res.status(500).send({error: e.message})
  }
});
module.exports = router;
