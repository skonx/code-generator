const express = require('express');
const morgan = require('morgan');

const blockchain = require('./blockchain');
blockchain.init();

const app = express();
app.use(morgan("dev"));

app.get("/", (req,res) => res.send(blockchain.map));

const port = process.env.PORT || 9000;
app.listen(port, () => console.log(`NodeJS server listening on port ${port}`));