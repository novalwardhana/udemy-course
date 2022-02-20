/* File System
  Merupakan standard libray untuk memanioulasi file system, ada 3 jenis library:
  a. Library yang bersifat blocking atau asynchronous
  b. Library yang bersifat non-blocking menggunakan callback
  c. Library yang bersifat non-blocking tapi menggunakan promise, disarankan menggunakan promise
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/fs.html
*/

/* syncronous */
import fs from "fs"

const bufferSync = fs.readFileSync("17-file-system-sync.mjs")
console.info("bufferSync toString: ", bufferSync.toString())

const bufferSync2 = fs.readFileSync("test-read-txt.txt").toString()
console.info(`bufferSync2: `, bufferSync2)

const bufferSync3 = fs.readFileSync("test-read-json.json").toString()
console.info(`bufferSync3: `, JSON.parse(bufferSync3))

fs.writeFileSync("test-write-sync.txt", "Example file 123456")

