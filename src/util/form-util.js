const formidable = require('formidable');

const parseForm = (req) => {
    var form = new formidable.IncomingForm();
    form.encoding = 'utf-8';

    return new Promise(function (resolve, reject) {
        form.parse(req, (err, fields, files) => {
            if (err) {
                reject(err);
            } else {
                resolve([fields, files]);
            }
        });
    });
}

module.exports = { parseForm }