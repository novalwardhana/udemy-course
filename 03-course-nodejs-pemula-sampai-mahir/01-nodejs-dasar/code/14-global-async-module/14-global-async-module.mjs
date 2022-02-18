/*
  GLobal Async Module
  a. Saat kita belajar Javascript async-await, biasanya kita perlu membuat terlebih dahulu function
    yang kita tandai sebagai async
  b. Saat kita menggunakan module, secara default global code adalah async, oleh karena itu kita bisa
    menggunakan async-await
  c. Kecuali jika kita membuat function, maka function tersebut harus kita tandai sebagai async jika 
    ingin menggunakan async-await
  d. Hanya bisa digunakan pada file dengan ekstensi .mjs dan nodeJS versi 16 keatas
*/

async function samplePromise() {
  return Promise.resolve(`Hello world`)
}
const name = await samplePromise()
console.info(name)