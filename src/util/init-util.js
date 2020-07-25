const fs = require('fs');
const path = require('path');
const { access, readFile } = require('fs').promises;
const { buildImageRecords } = require('./image-util');

const initData = async (db, dir = 'init') => {
    try {
        await access(path.join(dir, 'data.json'), fs.constants.R_OK)
    } catch (e) {
        return
    }
    const data = await readFile(path.join(dir, 'data.json'), 'utf-8')
    const json = JSON.parse(data);
    for (let record of json.records) {
        console.log(`\t${record.name}`);
        const records = await buildImageRecords(path.join(dir, record.image), record.profiles);
        db.add(record.name, records);
    }
}

module.exports = { initData }