const { validateProfiles } = require('../../src/util/size-util');

describe('test size utility', () => {
    describe('validateProfiles', () => {
        test('given list of valid profils should return invalid profiles', () => {
            const profiles = [{
                name: 'large',
                width: 1024,
                heigh: 680
            },{
                name: 'medium',
                width: 800,
                heigh: 600
            },{
                name: 'small',
                width: 240,
                heigh: 120
            }];
            expect(validateProfiles(profiles, 2000, 2000)).toHaveLength(0);
            expect(validateProfiles(profiles, 1000, 400)).toHaveLength(1);
            expect(validateProfiles(profiles, 100, 200)).toHaveLength(3);
        })

    })

})