/*
  Path
  Adalah standard library yang bisa kita gunakan untuk bekerja pada lokasi file atau direktori
  Refersnsi: https://nodejs.org/dist/latest-v16.x/docs/api/path.html
*/

import path from "path"

const filename = "/home/novalwardhana/temporary.txt"
const dirname = path.dirname(filename)
const basename = path.basename(filename)
const extname = path.extname(filename)
const parse = path.parse(filename)
console.info(`filename: ${filename}`)
console.info(`dirname: ${dirname}`)
console.info(`basename: ${basename}`)
console.info(`extname: ${extname}`)
console.info(`parse: `, parse)

console.info("\n ========== \n")

const filename2 = "/home/novalwardhana/document/photo/picture.jpeg"
const dirname2 = path.dirname(filename2)
const basename2 = path.basename(filename2)
const extname2 = path.extname(filename2)
const parse2 = path.parse(filename2)
console.info(`filename2: ${filename2}`)
console.info(`dirname2: ${dirname2}`)
console.info(`basename2: ${basename2}`)
console.info(`extname2: ${extname2}`)
console.info(`parse2: `, parse2)