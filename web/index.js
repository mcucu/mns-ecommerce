const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');

dotenv.config();
const appEnv = process.env;
const appPort = appEnv.APP_PORT || 8080;

const app = express();
app.use(cors());
app.use(express.static('public'))

/**
 * startServer
 */
server = app.listen(appPort, function() {
    console.info(`App is listening on http://localhos:${appPort}`);
});

/**
 * stopServer
 * @returns {void} - stop signal
 */
const stopServer = () => {
    server.close();
};

process.on('SIGINT', stopServer);
process.on('SIGTERM', stopServer);
