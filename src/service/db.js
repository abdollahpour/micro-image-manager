const { MongoClient, ObjectId } = require('mongodb');
const { ObjectID, Binary } = require('mongodb');
const { retryPromise } = require('../util/promise-util');
const fs = require('fs').promises;

class Db {

    constructor(mongodbUrl) {
        this.mongodbUrl = mongodbUrl;
    }

    async connect() {
        this.client = await retryPromise(
            () => MongoClient.connect(this.mongodbUrl, { useUnifiedTopology: true }),
            3000,
            10
        );
        this.db = this.client.db();
        this.images = this.db.collection('images');
        return this;
    }

    async disconnect() {
        await this.client.close();
    }

    /**
     * @param {{
     * image: string,
     * name: stirng, 
     * format: string, 
     * formatPriority: number,
     * profile: {width: number, height: number, name: string}
     * }} records 
     */
    async add(name, records) {
        const id = ObjectID();

        for (let record of records) {
            const buffer = await fs.readFile(record.image);

            await this.images.insertOne({
                id,
                name: name,
                data: Binary(buffer),
                created: new Date(),
                format: record.format,
                formatPriority: record.formatPriority,
                width: record.profile.width,
                height: record.profile.height,
                profile: record.profile.name,
                size: 35958
            });
        }

        // remove old records
        await this.images.deleteMany({ name, id: { $ne: id } });

        return id;
    }

    /**
     * 
     * @param {{
     * id: string, 
     * profile: string, 
     * formats: [string]
     * }} query 
     */
    async find(query) {
        const q = {
            format: { $in: query.formats }
        }
        if (query.id) q.id = ObjectID(query.id);
        if (query.name) q.name = query.name;
        if (query.profile) q.profile = query.profile;
        const sort = {
            width: -1,
            heigh: -1,
            formatPriority: 1
        }
        return (await this.images.find(q).limit(1).sort(sort).toArray())[0]
    }

}

module.exports = Db