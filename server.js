const http =require('http');
//const app = require('./app')
//const app = require('../app.js')


const app1 = require('./app')
const port  = process.env.PORT || 3001;

const server  = http.createServer(app1);

server.listen(port); 