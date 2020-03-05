const crypto = require('crypto');
const fs = require('fs');

const map = { blocks: {}, last_hash: 0 };
const cipher_algorithm = 'aes-192-cbc';
const salt = 'tR3ndEv_';

function generate_iv(code) {
    if (code.length >= 16) {
        return code.substring(0, 16);
    } else {
        let c = code;
        while (c.length < 16) {
            c += code;
        }
        console.warn("Secret code was too short for iv, new one generated : " + c.substring(0, 16));
        return c.substring(0, 16);
    }
}

function init_cipher(code) {
    const password = code;
    const key = crypto.scryptSync(password, salt, 24);
    const iv = generate_iv(code);
    return { key, iv };
}

function encrypt(value, secret) {
    const { key, iv } = init_cipher(secret.code);
    const cipher = crypto.createCipheriv(cipher_algorithm, key, iv);
    let encrypted = cipher.update(value, 'utf8', 'hex');
    encrypted += cipher.final('hex');
    return encrypted;
}

function unencrypt(encrypted, secret) {
    const { key, iv } = init_cipher(secret.code);
    const decipher = crypto.createDecipheriv(cipher_algorithm, key, iv);
    let decrypted = decipher.update(encrypted, 'hex', 'utf8');
    decrypted += decipher.final('utf8');
    return decrypted;
}

function hash_block(block) {
    const hash = crypto.createHash('sha512');
    return hash.update(JSON.stringify(block), 'utf-8').digest('hex');
}

function create_block(content) {
    const block = {
        content: content,
        //TODO : use process.hrtime() https://nodejs.org/docs/latest-v12.x/api/process.html#process_process_hrtime_time
        timestamp: new Date().getTime(),
        previous_hash: map.last_hash
    };
    return block;
}

function saveInChain(data) {
    console.log(`Saving data : ${data}`);

    const secret = JSON.parse(fs.readFileSync('/tmp/secret-code/secret.json'));
    const content = encrypt(data, secret);
    console.log('Data : \033[5;32mencrypted\033[0m');

    const block = create_block(content);
    const h = hash_block(block);
    console.log(`Block hash : ${h}`);

    map.last_hash = h;
    map.blocks[h] = block;

    console.log('Blockchain : \033[5;32mupdated\033[0m')

    return { secret, h }
}

function control_bc_integrity() {
    console.log("\033[5;33mControlling blockchain integrity...\033[0m");
    if (Object.keys(map.blocks).length > 1) {
        let next_hash = map.last_hash;
        let current_hash = map.blocks[map.last_hash].previous_hash;
        while (current_hash !== 0) {
            console.log(`controlling block integrity - hash : ${current_hash}`);
            const block = map.blocks[current_hash];
            const h = hash_block(block);
            if (h !== current_hash) {
                throw Error(`Corruption detected - the computed hash and the stored hash are different :\n${h}\n${current_hash}`);
            }

            const next_block = map.blocks[next_hash];

            if (h !== next_block.previous_hash) {
                throw Error(`Corruption detected - the computed hash and the previous hash stored in the next block are different :\n${h}\n${next_block.previous_hash}`);
            }

            const hash_next_block = hash_block({
                content: next_block.content,
                timestamp: next_block.timestamp,
                previous_hash: h
            });

            if (hash_next_block !== next_hash) {
                throw Error(`Corruption detected - the computed hash of the next block and the next block hash are different :\n${hash_next_block}\n${next_hash}`);
            }

            console.log("block " + current_hash + " : \033[5;32mOK\033[0m")
            current_hash = block.previous_hash;
            next_hash = next_block.previous_hash;
        }
    }
    console.log("Blockchain integrity: \033[1;32mOK\033[0m");
}

function init() {

    console.log("\033[5;33mBlockchain initialization...\033[0m");

    const n_seed = 5;

    const uncrypted = []

    // add n_seed blocks
    for (let i = 0; i < n_seed; i++) {
        uncrypted.push(saveInChain(JSON.stringify({ data: `seed-${i}` })));
    }

    // control the unencrypted content of the block
    if (uncrypted
        .filter(({ secret, h }, i) => JSON.stringify({ data: `seed-${i}` }) === unencrypt(map.blocks[h].content, secret))
        .length === n_seed) {
        control_bc_integrity();
        console.log("Blockchain : \033[1;32minitialized\033[0m");
    } else {
        console.log("Blockchain : \033[1;31mnot initialized\033[0m");
        throw Error("Encryption error occurs during blockchain initialization");
    }
}

module.exports = {
    map,
    init
};