const fs = require('fs');
const { identifyImage } = require('../util/http-util');
const { ObjectID } = require('mongodb');

module.exports = async (req, res, db) => {
    const { image, profile, formats } = identifyImage(req)
    const loaded = await db.find({ id: ObjectID(image), profile, formats });

    if (loaded) {
        res.writeHead(200, {'Content-Type': `image/${loaded.format}` });
        res.end(loaded.data.buffer);
    } else {
        const notFound = profile ? `${__dirname}/../../assets/${profile}.png` : undefined;
        const readStream = fs.createReadStream(notFound);

        readStream.on('open', () => {
            res.writeHead(404, {
                'content-type': 'image/png'
            });
            readStream.pipe(res);
        });

        readStream.on('error', () => {
            res.writeHead(404);
            res.end();
        });
    }
}