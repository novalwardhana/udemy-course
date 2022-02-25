/*
  HTTP Server
  a. Standard library HTTP selain bisa sebagai HTTP CLient, bisa juga sebagaiu HTTP Server
  b. Untuk kasus sederhana, cocok sekali jika ingin menggunakan HTTP Server menggunakan standard library NodeJS,
    namun untuk kasus yang lebih kompleks, direkomndasikan menggunakan library
    atau framework yang lebih mudah penggunaannya
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/http.html
    https://nodejs.org/dist/latest-v16.x/docs/api/https.html
*/

import http from "http"

const server = http.createServer((request, response) => {
  console.info(request.method)
  console.info(request.url)
  //console.table(request.headers)

  if (request.method === "POST") {
    request.addListener("data", (data) => {
      response.setHeader("Content-Type", "application/json")
      response.write(data)
      response.end()
    })
  } else {
    if (request.url === "/noval") {
      response.write("Hello Noval")
    } else {
      response.write("Welcome to my API")
    }
    response.end()
  }
})

server.listen("3000")