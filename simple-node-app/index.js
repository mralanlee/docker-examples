const http = require('http');

const PORT = process.env.PORT;
const ENV = process.env.NODE_ENV;

console.log(`Listening on http://localhost:${PORT}`);

http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.write(`Hello World! From Enviroment: ${ENV}, Port: ${PORT}`);
  res.end();
}).listen(PORT);
