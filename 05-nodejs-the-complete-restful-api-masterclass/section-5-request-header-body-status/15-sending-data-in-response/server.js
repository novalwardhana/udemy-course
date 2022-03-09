const http = require("http")

const data = [
  {
    id: 1,
    name: "Ronaldo"
  },
  {
    id: 2,
    name: "Messi"
  },
  {
    id: 3,
    name: "Lewandowski"
  }
]
const server = http.createServer((req, res) => {
  res.end(JSON.stringify({
    success: true,
    results: data.length,
    data: data
  }))
})
server.listen("3000", () => {
  console.info("Run server")
})