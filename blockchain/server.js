const express = require('express');
const morgan = require('morgan');

const blockchain = require('./blockchain');
blockchain.init();

const app = express();
app.use(morgan("dev"));

app.get("/", (req, res) => res.send(blockchain.map));

const port = process.env.PORT || 9000;
app.listen(port, () => console.log("\033[1;34mNodeJS server listening on port " + port + "\033[0m"));