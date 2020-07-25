const { initData } = require('../../src/util/init-util');
const Db = require('../../src/service/db');

describe('test init utility', () => {
    const db = new Db('mongodb://localhost:27017/image-manager');

    beforeAll(async () => await db.connect());
    afterAll(async () => await db.disconnect());

    describe('initData', () => {
        test('given test import dir should initilize the database with images', async() => {
            await initData(db, 'test/resources')
        });
    })

})