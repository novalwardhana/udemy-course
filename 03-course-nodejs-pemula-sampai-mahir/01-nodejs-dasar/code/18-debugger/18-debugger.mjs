/*
  Debugger
  1. Jika kita menjalankan file Javascript hanya menggunakan perintah node namafile.js, maka secara default tidak akan ke mode debug
  2. Untuk menjalankan mode debug, kita perlu menambahkan perintah: node inspect namafile.js
  3. Perintah debugger
    a. cont, c: continue
    b. next: step code
    c. step: masuk ke function
    d. out, o: step out
    e. pause: pause running code
  Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/debugger.html
*/

/* Cara menjalankan:
  node inspect 18-debugger.mjs
  c
  watch("name")
  c
*/

function sayHello(name) {
  debugger
  return `Hello ${name}`
}
debugger
const name = "noval"
console.info(sayHello(name))

const squareNumber = (...datas) => {
  debugger
  const result = datas.map(element => Math.pow(element, 2))
  return result
}
debugger
const numbers = [5, 4, 8, 9, 12, 18]
console.info(squareNumber(...numbers))