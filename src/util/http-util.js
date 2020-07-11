const url = require('url');

const webpSupported = (req) =>
    req.headers['accept']?.indexOf('image/webp') > -1 ? 'webp' : '';

const identifyImage = (req) => {
    const requestUrl = req.url.toLowerCase();
    const results = requestUrl.match(/\/image\/([0-9a-z]{24})(\.(png|gif|jpg|jpeg|webp))?(\?.*)?/);

    if (results) {
        const { query } = url.parse(requestUrl, true);
        const formats = [...new Set([].concat(
            query['format']?.replace('jpg', 'jpeg').split(',')).concat([results[3   ], webpSupported(req), 'jpeg', 'png']
            ))].filter(i => i);
        const profile = query.profile

        return { image: results[1], profile, formats };
    }
    throw Error(`Illegal path: ${req.url}`);
}

module.exports = { identifyImage }