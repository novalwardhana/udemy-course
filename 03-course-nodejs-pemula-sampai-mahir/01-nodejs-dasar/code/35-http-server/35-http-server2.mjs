import http from "http"

const server = http.createServer((request, response) => {
  if (request.method === "POST" && request.url === "/post-data") {
    request.addListener("data", (data) => {
      response.write(data)
      response.end()
    })
  } else {
    response.write("Welcome to my API")
    response.end()
  }
})
server.listen("3000")