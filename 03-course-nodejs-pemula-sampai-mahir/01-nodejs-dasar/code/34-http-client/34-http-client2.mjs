import http from "http"

const client = http.request("http://localhost:3000/post-data", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Accept": "json"
  }
}, (response) => {
  response.addListener("data", (data) => {
    console.info(`Response data: `, data.toString())
  })
})
const body = JSON.stringify(
  {
    name: "Ronaldo",
    number: 7,
    club: "Manchester United"
  }
)
client.write(body)
client.end()

const client2 = http.request("http://localhost:3000", {
  method: "GET",
  headers: {
    "content-Type": "application/json",
    "Accept": "json"
  }
}, (response) => {
  response.addListener("data", (data) => {
    console.info(`Response data: `, data.toString())
  })
})
client2.end()