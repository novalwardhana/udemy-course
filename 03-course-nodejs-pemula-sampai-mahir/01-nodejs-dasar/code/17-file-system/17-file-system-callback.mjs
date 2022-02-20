/* File System
  Merupakan standard libray untuk memanioulasi file system, ada 3 jenis library:
  a. Library yang bersifat blocking atau asynchronous
  b. Library yang bersifat non-blocking menggunakan callback
  c. Library yang bersifat non-blocking tapi menggunakan promise, disarankan menggunakan promise
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/fs.html
*/

import fs from "fs"

await fs.readFile("17-file-system-callback.mjs", (error, response) => {
  if (error !== null) {
    console.info(`bufferCallback error: `, error)
    return
  }
  console.info(`bufferCallback response: `, response.toString())
})

await fs.readFile("test-read-txt.txt", (error, response) => {
  if (error !== null) {
    console.info(`bufferCallback2 error: `, error)
    return
  }
  console.info(`BufferCallback2 response: `, response.toString())
})

await fs.readFile("test-read-json.json", (error, response) => {
  if (error !== null) {
    console.info(`bufferCallback3 error: `, error)
    return
  }
  console.info(JSON.parse(response.toString()))
})

const fileDataTxt = "ABCDEFG 123456"
await fs.writeFile("test-write-callback.txt", fileDataTxt, (error) => {
  if (error !== null) {
    console.info(`Error write file: ${error.message}`)
    return
  }
  console.info(`Success write file`)
})

const fileDataJSON = [
  {"name": "Sate Ayam", "price": 15000},
  {"name": "Sate Kambing", "price": 25000}
]
await fs.writeFileSync("test-write-callback.json", JSON.stringify(fileDataJSON), (error) => {
  if (error !== null) {
    console.info(`Error write file: ${error.message}`)
    return
  }
  console.info(`Success write file json`)
})
