const crypto = require('crypto');
const fs = require('fs');

const map = {};
const cipher_algorithm = 'aes-192-cbc';
const salt = 'tr3ndEv_';

function generate_iv(secret) {
    if (secret.code.length >= 16) {
        return secret.code.substring(0, 16);
    } else {
        let c = secret.code;
        while (c.length < 16) {
            c += secret.code;
        }
        console.warn("Secret code was too short for iv, new one generated : " + c.substring(0, 16));
        return c.substring(0, 16);
    }
}

function init_cipher(secret) {
    const password = secret.code;
    const key = crypto.scryptSync(password, salt, 24);
    const iv = generate_iv(secret);
    return { key, iv };
}

function encrypt(value, secret) {
    const { key, iv } = init_cipher(secret);
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
}

module.exports = {
    map,
    init
};