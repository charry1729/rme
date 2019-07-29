const express = require('express');

const router = express.Router();

const read_file = require('../roh/controller.js')

//router.get('/:Id',read_file.get_data);

router.get('/query',read_file.get_data);

router.post('/post', read_file.load_data);

module.exports = router;
