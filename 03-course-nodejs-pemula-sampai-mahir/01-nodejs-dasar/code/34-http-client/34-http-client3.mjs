import http from "http"

const client = http.request("http://localhost:3000", {
  method: "GET",
  headers: {
    "Content-Type": "application/json",
    "Accept": "json"
  }
}, (response) => {
  response.addListener("data", (data) => {
    console.info(`Receiving response data: ${data.toString()}`)
  })
})
client.end()

const clientStat = http.request("http://localhost:3000/statistic", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Accept": "json"
  }
}, (response) => {
  if (response.statusCode === 200) {
    response.addListener("data", (data) => {
      console.info(`Receiving response data: `, JSON.parse(data.toString()))
    })
  }
})
clientStat.write(JSON.stringify([10, 8, 6, 4, 10, 7, 8, 9, 7, 9]))
clientStat.end()

const clientSquare = http.request("http:localhost:3000/square", {
  method: "POST",
  headers: {
    "Contenbt-Type": "application/json",
    "Accept": "json"
  }
}, (response) => {
  if (response.statusCode === 200) {
    response.addListener("data", (data) => {
      console.info(`Receiving response data: `, JSON.parse(data.toString()))
    })
  }
})
clientSquare.write(JSON.stringify([10, 8, 6, 4, 10, 7, 8, 9, 7, 9]))
clientSquare.end()