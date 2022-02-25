/*
  HTTP Client
  a. HTTP Client: melakukan simulasi HTTP request menggunakan NodeJS
  b. Ada 2 jenis HTTP di NodeJS yaitu HTTP dan HTTPS
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/http.html
    https://nodejs.org/dist/latest-v16.x/docs/api/https.html
*/

import { request } from "http"
import https from "https"

const url = "https://hookb.in/E7GEJlan1RiVjY669p31"
const request1 = https.request(url, {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Accept": "application/json"
  }
}, (response) => {
  response.addListener("data", (data) => {
    console.info(`Receiving data ${data}`)
  })
})

const body = {
  firstName: "Noval",
  lastName: "Wardhana"
}
request1.write(JSON.stringify(body))
request1.end()