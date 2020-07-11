const Db = require('../../src/service/db');

describe('test db service', () => {
    const db = new Db('mongodb://localhost:27017/image-manager');

    beforeAll(async () => await db.connect());
    afterAll(async () => await db.disconnect());

    describe('add image', () => {
        test('given image file and profiles should add images to database', async () => {
            const id = await db.add([{
                image: `${__dirname}/../resources/1280x1080.png`,
                format: 'png',
                formatPriority: 1,
                profile: {
                    name: 'large',
                    width: 1920,
                    height: 1080
                }
            }]);
            expect(id).not.toBeNull();
        });
    })

    describe('get image', () => {
        test('given query should find the image', async () => {
            const id = await db.add([{
                image: `${__dirname}/../resources/1280x1080.png`,
                format: 'png',
                formatPriority: 1,
                profile: {
                    name: 'large',
                    width: 1920,
                    height: 1080
                }
            }]);

            const image1 = await db.find({id, profile: 'large', formats: ['png']});
            expect(image1.format).toBe('png')

            const image2 = await db.find({id, formats: ['webp', 'png']});
            expect(image2.format).toBe('png')

            const image3 = await db.find({id, profile: 'large', formats: ['webp']});
            expect(image3).toBeUndefined()
        });
    })

})