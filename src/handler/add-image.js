const sharp = require('sharp')
const { parseForm } = require('../util/form-util');
const { loadProfiles, validateProfiles } = require('../util/size-util');
const { buildImageRecords } = require('../util/image-util');

module.exports = async (req, res, db) => {
    const [fields, files] = await parseForm(req);
    const profiles = loadProfiles(fields);
    const name = fields.name;
    if (!name) {
        res.writeHead(400);
        return res.end(JSON.stringify({
            status: 400,
            error: 'Name is required'
        }));
    }
    if (name.match('[a-zA-Z0-9]{24}')) {
        res.writeHead(400);
        return res.end(JSON.stringify({
            status: 400,
            error: 'Name cannot be with mongodb ObjectID hex format'
        }));
    }

    if (Object.keys(files).length !== 1) {
        res.writeHead(400);
        return res.end(JSON.stringify({
            status: 400,
            error: 'You need one image file'
        }));
    }
    const imageFile = files[Object.keys(files)[0]].path;

    const sharped = await sharp(imageFile);
    const metadata = await sharped.metadata();

    if (fields['enlarge']?.toLowerCase() !== 'true') {
        const validation = validateProfiles(profiles, metadata.width, metadata.height);
        if (validation.length > 0) {
            res.writeHead(400);
            return res.end(JSON.stringify({
                status: 400,
                error: 'Invalid image sizes',
                errors: validation
            }));
        }
    }

    const records = await buildImageRecords(imageFile, profiles);
    const id = await db.add(name, records);

    res.writeHead(200);
    res.end(JSON.stringify({
        id,
        profiles
    }));
}