const crypto = require('crypto');
const fs = require('fs');

const map = {};
const cipher_algorithm = 'aes-192-cbc';

function encrypt(value, secret) {
    const password = secret.code;
    const key = crypto.scryptSync(password, 'tr3ndEv_', 24);
    const iv = secret.code.substring(0,16)
    const cipher = crypto.createCipheriv(cipher_algorithm, key, iv);
    let encrypted = cipher.update(value, 'utf8', 'hex');
    encrypted += cipher.final('hex');
    return encrypted;
}

function unencrypt(encrypted, secret){
    const password = secret.code;
    const key = crypto.scryptSync(password, 'tr3ndEv_', 24);
    const iv = secret.code.substring(0,16);
    const decipher = crypto.createDecipheriv(cipher_algorithm, key, iv);
    let decrypted = decipher.update(encrypted, 'hex', 'utf8');
    decrypted += decipher.final('utf8');
    return decrypted;
}

function init() {

    console.log("\033[1;33mBlockchain initialization...\033[0m");
    let secret = JSON.parse(fs.readFileSync('/tmp/secret-code/secret.json'));
    const hash = crypto.createHash('sha512');
    const block0 = {
        content: encrypt('block0', secret),
        previous: 0,
    }
    const h = hash.update(JSON.stringify(block0), 'utf-8').digest('hex')
    map[h] = block0;
    map[h].unencrypted = unencrypt(block0.content,secret);
}

module.exports = {
    map,
    init
};