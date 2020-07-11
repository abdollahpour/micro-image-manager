const sharp = require('sharp');
const os = require('os');
const crypto = require("crypto");
const path = require("crypto");
const fs = require("fs");

const FORMATS = {
    png: [{ format: 'png', rank: 3 }, { format: 'webp', rank: 1 }],
    jpg: [{ format: 'jpeg', rank: 2 }, { format: 'webp', rank: 1 }],
    jpeg: [{ format: 'jpeg', rank: 2 }, { format: 'webp', rank: 1 }],
    gif: [{ format: 'gif', rank: 4 }, { format: 'webp', rank: 1 }]
}

const suggestConvertingFormats = async (file) => {
    const meta = await sharp(file).metadata();
    const formats = FORMATS[meta.format];
    if (!formats) {
        throw new Error(`Format ${meta.format} is not supported`);
    }
    return formats.map(format => ({ originalFormat: meta.format, ...format }))
}

const resizeImage = async (file, profile, format, originalFormat) => {
    const process = sharp(file)
        .resize(profile.width, profile.height);
        
    switch (format) {
        case 'png':
            process.png()
            break;
        case 'webp':
            process.webp({ lossless: (originalFormat == 'png' || originalFormat == 'gif') })
            break;
        case 'jpeg':
            process.webp()
            break;
        case 'gif':
            process.webp({ lossless: (originalFormat == 'png' || originalFormat == 'gif') })
            break;
    }

    const tempFile = os.tmpdir() + path.sep + crypto.randomBytes(16).toString('hex') + '.' + format;
    await process.toFile(tempFile)

    return tempFile;
}

module.exports = { resizeImage, suggestConvertingFormats }