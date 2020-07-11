const http = require('http');
const { promisify } = require('util');
const getImage = require('./handler/get-image');
const addImage = require('./handler/add-image');
const Db = require('./service/db');

const mongodbUrl = process.env['MONGO_URL'] ?? 'mongodb://localhost:27017/image-manager';
const httpPort = process.env['HTTP_PORT'] ?? 8080;

const requestListener = (db) => 
    async (req, res) => {
        if (/\/image\/[0-9a-z_]+/.test(req.url)) {
            getImage(req, res, db);
        } else if (req.url === '/api/v1/images') {
            addImage(req, res, db);
        } else {
            res.writeHead(404);
            return res.end(JSON.stringify({
                status: 404,
                error: `'${req.url}' does not exist`
            }));
        }
    }

const main = async () => {
    const db = new Db(mongodbUrl);
    const server = http.createServer(requestListener(await db.connect()));
    await server.listen(httpPort)
    console.log(`Server ready at ${httpPort}`);
}

main();