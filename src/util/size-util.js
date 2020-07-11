const loadProfiles = (fields) => Object.keys(fields)
    .filter(key => /profile_[a-z]+/.test(key))
    .map(key => [key.substring(8), fields[key]])
    .filter(([_, value]) => (/[0-9]+x[0-9]+/.test(value)))
    .map(([name, value]) => {
        const parts = value.split('x');
        return { name, width: parseInt(parts[0]), height: parseInt(parts[1]) }
    });

const validateProfiles = (profiles, width, height) => profiles
    .filter(profile => profile.width > width || profile.height > height)
    .map(profile => {
        message: `image size is not valid. Width (${profile.width}) has to be larger than ${width} and heigh has to be large than ${height}`,
        width, height, profile
    })

module.exports = { loadProfiles, validateProfiles }