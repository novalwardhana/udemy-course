/*
  GLobal Async Module
  a. Saat kita belajar Javascript async-await, biasanya kita perlu membuat terlebih dahulu function
    yang kita tandai sebagai async
  b. Saat kita menggunakan module, secara default global code adalah async, oleh karena itu kita bisa
    menggunakan async-await
  c. Kecuali jika kita membuat function, maka function tersebut harus kita tandai sebagai async jika 
    ingin menggunakan async-await
  d. Hanya bisa digunakan pada file dengan ekstensi .mjs
*/

/* Pada Javascript */
function samplePromise() {
  return Promise.resolve("Hallo Noval")
}
async function run() {
  const name = await samplePromise()
  console.info(name)
}
run()

const squareValue = (...datas) => {
  return new Promise((resolve, reject) => {
    const result = datas.map(element => Math.pow(element, 2))
    if (result) {
      resolve(result)
    } else {
      reject(`Not valid`)
    }
  })
}
const runSquareValue = async function() {
  const result = await squareValue(8, 9, 4, 5, 7, 6, 9, 8, 9, 8)
  return result
}
runSquareValue()
  .then((response) => {
    console.info(response)
  })