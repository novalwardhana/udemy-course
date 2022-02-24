/*
  Zlib
  a. Zlib adalah standard library untuk melakukan kompresi menggunakan Gzip
  b. Kegunaan: untuk penyimpanan file agar ukuran tidak terlalu besar
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/zlib.html
*/

import fs from "fs"
import zlib from "zlib"

const file = fs.readFileSync("31-zlib-compress.mjs")
const fileCompress = zlib.gzipSync(file)
fs.writeFileSync("compress-file.txt", fileCompress)

const file2 = fs.readFileSync("31-zlib-decompress.mjs")
const fileCompress2 = zlib.gzipSync(file2)
fs.writeFileSync("compress-file2.txt", fileCompress2)