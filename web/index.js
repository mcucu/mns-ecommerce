const express = require('express');
const cors = require('cors');

const app = express();
app.use(cors());
app.use(express.static('public'))

/**
 * startServer
 */
server = app.listen(8080, function() {
    console.info(`App is listening on http://localhost:8080`);
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
