const express = require('express');
const router = express.Router();
const applications = require('../controllers/applications');


router.get('/', async (req, res) => {
    try{
        // const {data} = await applications.getAll(req.query);
        const {data} = await applications.getAll(req.query);
        res.status(200).send({
            results: data.Records,
            totalCount: data.Count
        });
    }catch (e) {
        res.status(500).send({error: e.message})
    }
});
module.exports = router;
