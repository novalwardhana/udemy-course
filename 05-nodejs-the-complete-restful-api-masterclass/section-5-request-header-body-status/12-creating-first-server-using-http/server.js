const http = require("http")

const server = http.createServer((request, response) => {
  response.end("First server")
})
server.listen(3000, () => {
  console.info("Run localhost")
})