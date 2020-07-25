const { identifyImage } = require('../../src/util/http-util');

describe('test image utility', () => {
    describe('identifyImage', () => {
        test('given http request with valid pass should return image properties', () => {
            // With webp
            expect(identifyImage({
                url: '/image/name.jpeg',
                headers: {
                    accept: 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9'
                }
            })).toEqual({
                name: 'name',
                formats: ['jpeg', 'webp', 'png']
            });

            // Without webp
            expect(identifyImage({
                url: '/image/560c24b853b558856ef193a3.jpeg',
                headers: {
                    accept: 'text/html,application/xhtml+xml,application/xml;q=0.9,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9'
                }
            })).toEqual({
                id: '560c24b853b558856ef193a3',
                formats: ['jpeg', 'png']
            });
        });
    })

})