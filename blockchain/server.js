'use strict';

const express = require('express');
const morgan = require('morgan');

const app = express();
app.use(morgan("dev"));

const blockchain = []

app.get("/", (req,res) => res.send(blockchain));

const port = process.env.PORT || 9000;
app.listen(port, () => console.log(`NodeJS server listening on port ${port}`));