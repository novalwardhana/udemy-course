/*
  Buffer
  a. Buffer merupakan object yang berisikan urutan byte dengan panjang yang tetap
  b. Buffer merupakan turunan dari tipe data Uint8Array
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/buffer.html

  Buffer Encoding
  a. Buffer juga bisa digunakan untuk melakukan encoding dari satu encoding ke encoding yang lain
  b. Ada banyak encoding yang didukung oleh buffer, misal utf8, ascii, hex, base64
*/

const buffer = Buffer.from("Noval Kusuma Wardhana")
console.info(`buffer: `, buffer)
console.info(`buffer toString: `, buffer.toString())
buffer.reverse()
console.info(`buffer toString: `, buffer.toString())

const buffer2 = Buffer.from("Manchester United")
console.info(`buffer2: `, buffer2)
console.info(`buffer2 toString: `, buffer2.toString())
buffer2.reverse()
console.info(`Buffer2 toString: `, buffer2.toString())

const buffer3 = Buffer.from("Novalita Kusuma Wardhana", "utf8")
const buffer3toString = buffer3.toString()
const buffer3Base64 = buffer3.toString("base64")
const buffer3Hex = buffer3.toString("hex")
console.info(`buffer3toString: ${buffer3toString}`)
console.info(`buffer3Base64: ${buffer3Base64}`)
console.info(`buffer3Hex: ${buffer3Hex}`)

const buffer4 = Buffer.from("Real Madrid")
const buffer4Base64 = buffer4.toString("base64")
const buffer4Hex = buffer4.toString("hex")
console.info(`buffer4: ${buffer4.toString()}`)
console.info(`buffer4Base64: ${buffer4Base64}`)
console.info(`buffer4Hex: ${buffer4Hex}`)

const buffer5 = Buffer.from("Tm92YWxpdGEgS3VzdW1hIFdhcmRoYW5h", "base64")
const buffer5toString = buffer5.toString("base64")
const buffer5Utf8 = buffer5.toString("utf8")
const buffer5Hex = buffer5.toString("hex")
console.info(`buffer5toString: ${buffer5toString}`)
console.info(`buffer5Utf8: ${buffer5Utf8}`)
console.info(`buffer5Hex: ${buffer5Hex}`)

const buffer6 = Buffer.from("5265616c204d6164726964", "hex")
const buffer6toString = buffer6.toString("hex")
const buffer6Base64 = buffer6.toString("base64")
const buffer6Utf8 = buffer6.toString("utf8")
console.info(`buffer6toString: ${buffer6toString}`)
console.info(`buffer6Base64: ${buffer6Base64}`)
console.info(`buffer6Utf8: ${buffer6Utf8}`)

const buffer7 = Buffer.byteLength("4e6f76616c697461204b7573756d61205761726468616e61")
console.info(`buffer7: ${buffer7}`)

const buffer8 = Buffer.byteLength("Novalita")
console.info(`buffer8: ${buffer8}`)
