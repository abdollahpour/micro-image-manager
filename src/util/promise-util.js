const wait = ms => new Promise(r => setTimeout(r, ms));

const retryPromise = (operation, delay, times) => new Promise((resolve, reject) => {
    return operation()
        .then(resolve)
        .catch((reason) => {
            if (times - 1 > 0) {
                console.log("Retry to connect to mongodb...")
                return wait(delay)
                    .then(retryPromise.bind(null, operation, delay, times - 1))
                    .then(resolve)
                    .catch(reject);
            }
            return reject(reason);
        });
});

module.exports = { retryPromise };