import { Logger } from 'logger';

import * as http from 'http';

const hostname = '127.0.0.1';
const port = 3000;

const logger = new Logger();
logger.info('hello nolleh');

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('Hello World');
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
