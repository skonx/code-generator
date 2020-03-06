const express = require('express');
const morgan = require('morgan');
const bodyParser = require('body-parser');
const blockchain = require('./blockchain');
const { NS_PER_SEC } = require('./nano');

blockchain.init();

const app = express();
app.use(morgan("dev"));
app.use(bodyParser.json());
app.get("/content", (req, res) => {
    const hash = req.header("Block-Hash");
    const code = req.header("Secret-Code");

    if (!hash || !code) {
        res.status(401).send();
    } else {
        const block = blockchain.map.blocks[hash];

        if (!block) {
            res.status(404).json({
                status: 'error',
                message: `Block ${hash} not found`
            });
        } else {
            try {
                const time = process.hrtime();
                const decrypted_content = blockchain.decrypt(block.content, { code: code });
                const diff = process.hrtime(time);
                res.header("Decryption-Time-ns", `${diff[0] * NS_PER_SEC + diff[1]}`);
                res.json(JSON.parse(decrypted_content));
            } catch (error) {
                console.warn(error);
                res.status(401).json({
                    status: 'error',
                    message: `Block ${hash} cannot be decrypted`,
                });
            }
        }
    }
});

app.get("/control-integrity",(req,res) => {
    try{
        blockchain.control_integrity();
        res.send();
    }catch (error){
        console.error(error);
        res.json({
            status:'error',
            message: 'Blockchain integrity : corrupted'
        });
    }
});

app.get("/", (req, res) => res.send(blockchain.map));

app.post("/save", (req, res) => {
    const body = req.body;
    try {
        const { secret, h: hash} = blockchain.update(body);
        res.status(201).json({ secret, hash });
    } catch (error) {
        console.log(error);
        res.status(500).json({
            status: 'error',
            message: 'content cannot be saved in blockchain'
        });
        process.exit(1);
    }
});

const port = process.env.PORT || 9000;
app.listen(port, () => console.log("\033[1;34mNodeJS server listening on port " + port + "\033[0m"));