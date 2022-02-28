import http from "http"

const server = http.createServer((request, response) => {
  response.write("Todolist API")
  response.end()
})
server.listen(3030)