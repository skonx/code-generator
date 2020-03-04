'use strict';

const express = require('express');
const morgan = require('morgan');
const fs = require('fs');
var crypto = require('crypto');

const app = express();
app.use(morgan("dev"));

//use map instead 
const blockchain = [];

JSON.parse(fs.readFileSync('/tmp/secret-code/secret.json'))

app.get("/", (req,res) => res.send(blockchain));

const port = process.env.PORT || 9000;
app.listen(port, () => console.log(`NodeJS server listening on port ${port}`));