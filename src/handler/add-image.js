const sharp = require('sharp')
const { parseForm } = require('../util/form-util');
const { loadProfiles, validateProfiles } = require('../util/size-util');
const { resizeImage, suggestConvertingFormats } = require('../util/image-util');

module.exports = async (req, res, db) => {
    const [fields, files] = await parseForm(req);
    const profiles = loadProfiles(fields);

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

    const records = [];
    const suggestedFormats = await suggestConvertingFormats(imageFile);
    for (let profile of profiles) {
        for (let suggestedFormat of suggestedFormats) {
            const resizedImage = await resizeImage(imageFile, profile, suggestedFormat.format, suggestedFormat.originalFormat);
            records.push({
                image: resizedImage,
                profile,
                format: suggestedFormat.format,
                originalFormat: suggestedFormat.originalFormat,
                formatPriority: suggestedFormat.formatPriority
            });
        }
    }

    const id = await db.add(records);

    res.writeHead(200);
    res.end(JSON.stringify({
        id,
        profiles
    }));
}