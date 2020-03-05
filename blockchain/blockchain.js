const crypto = require('crypto');
const fs = require('fs');

const map = { last_hash: 0, blocks: {} };
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

function encrypt(value) {
    const secret = JSON.parse(fs.readFileSync('/tmp/secret-code/secret.json'));
    const { key, iv } = init_cipher(secret.code);
    const cipher = crypto.createCipheriv(cipher_algorithm, key, iv);
    let encrypted = cipher.update(value, 'utf8', 'hex');
    encrypted += cipher.final('hex');
    return encrypted;
}

function unencrypt(encrypted, secret) {
    const { key, iv } = init_cipher(secret);
    const decipher = crypto.createDecipheriv(cipher_algorithm, key, iv);
    let decrypted = decipher.update(encrypted, 'hex', 'utf8');
    decrypted += decipher.final('utf8');
    return decrypted;
}

function hash_block(block) {
    const hash = crypto.createHash('sha512');
    return hash.update(JSON.stringify(block), 'utf-8').digest('hex');
}

function create_block(data) {
    const block = {
        content: encrypt(data),
        timestamp: new Date().getTime(),
        previous_hash: map.last_hash
    };
    return block;
}

function saveInChain(data) {
    console.log(`Saving data : ${data}`);
    const block = create_block(data);
    const h = hash_block(block);
    map.last_hash = h;
    map.blocks[h] = block;
}

function init() {

    console.log("\033[5;33mBlockchain initialization...\033[0m");

    saveInChain('block-0');

    for (let i = 1; i <= 5; i++) {
        saveInChain(`block-${i}`);
    }

    console.log("\033[1;33mBlockchain initialized !\033[0m");
}

module.exports = {
    map,
    init
};