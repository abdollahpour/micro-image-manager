const url = require('url');

const webpSupported = (req) =>
    req.headers['accept']?.indexOf('image/webp') > -1 ? 'webp' : '';

const identifyImage = (req) => {
    const requestUrl = req.url.toLowerCase();
    const results = requestUrl.match(/\/image\/([0-9a-zA-Z_-]{1,128})(\.(png|gif|jpg|jpeg|webp))?(\?.*)?/);

    if (results) {
        const { query } = url.parse(requestUrl, true);
        const formatInQuery = query['format']?.replace('jpg', 'jpeg').split(',');
        const formatInName = results[3];
        const formatInHeader = webpSupported(req);
        const formatDefaults = ['jpeg', 'png'];
        const formats =  [...new Set(
            [].concat.apply([], [formatInQuery, formatInName, formatInHeader, formatDefaults]).filter(i => i)
        )];
        const profile = query.profile;
        return { image: results[1], profile, formats };
    }
    throw Error(`Illegal path: ${req.url}`);
}

module.exports = { identifyImage }