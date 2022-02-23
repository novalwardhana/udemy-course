/* 
  Stream
  a. Stream adalah standard library untuk kontrak aliran data di NodeJS
  b. Ada banyak sekali Stream object di NodeJS
  c. Stream bisa jadi object yang bisa dibaca atau bisa ditulis
    Stream turunan dari EventEmitter
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/stream.html
*/

import fs from "fs"

const writer = fs.createWriteStream("target.log")
writer.write("Noval\n")
writer.write("Kusuma\n")
writer.write("Wardhana\n")
writer.end()

const reader = fs.createReadStream("target.log")
reader.addListener("data", (buffer) => {
  console.info(buffer.toString())
})

const writer2 = fs.createWriteStream("target2.log")
const datas = ["Manchester United", "Chelsea", "Arsenal", "Liverpool", "Tottenham"]
for (const index in datas) {
  writer2.write(datas[index]+"\n")
}
writer2.end()
const reader2 = fs.createReadStream("target2.log")
reader2.addListener("data", (buffer) => {
  console.info(buffer.toString())
})