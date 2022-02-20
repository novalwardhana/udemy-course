/* File System
  Merupakan standard libray untuk memanioulasi file system, ada 3 jenis library:
  a. Library yang bersifat blocking atau asynchronous
  b. Library yang bersifat non-blocking menggunakan callback
  c. Library yang bersifat non-blocking tapi menggunakan promise, disarankan menggunakan promise
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/fs.html
*/

import fs from "fs/promises"

const bufferPromise = await fs.readFile("17-file-system-promise.mjs")
console.info(`bufferPromise: `, bufferPromise.toString())

const bufferPromise2 = await fs.readFile("test-read-txt.txt")
console.info(`bufferPromise2: `, bufferPromise2.toString())

const bufferPromise3 = await fs.readFile("test-read-json.json")
console.info(`bufferPromise3: `, JSON.parse(bufferPromise3.toString()))

const dataFileTxt = "ABCDEFG 123456"
await fs.writeFile("test-write-promise.txt", dataFileTxt)

const dataFileJSON = [
  {"name": "Sate Ayam", "price": 15000},
  {"name": "Sate Kambing", "price": 25000}
]
await fs.writeFile("test-write-promise.json", JSON.stringify(dataFileJSON))