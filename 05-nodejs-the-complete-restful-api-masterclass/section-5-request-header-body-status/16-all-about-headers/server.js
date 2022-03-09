const http = require("http")

const data = [
  {
    id: 1,
    name: "Manchester United"
  },
  {
    id: 2,
    name: "Chelsea"
  },
  {
    id: 3,
    name: "Arsenal"
  }
]
const server = http.createServer((req, res) => {
  res.setHeader("Content-Type", "application/json")
  res.setHeader("Content-Language", "en-US")
  res.setHeader("Date", new Date(Date.now()))
  res.setHeader("X-Powered-By", "node.js")
  res.setHeader("Author", "Novalwardhana")
  res.end(
    JSON.stringify({
      success: true,
      results: data.length,
      data: data
    })
  )
})
server.listen("3000", () => {
  console.info("Run server")
})