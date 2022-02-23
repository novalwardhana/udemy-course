/*
  Timer
  a. Timer merupakan standard library untuk melakukan scheduling
  b. Function di timer terdapat di globals, sehingga bisa menggunakannya tanpa melakukan import, namun semua menggunakan callback
  c. Jika ingin menggunakan versi Timer versi promise, kita bisa meng-import dari module timer/promise
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/timers.html
*/

setInterval(() => {
  console.info(`start interval date time at: ${new Date()}`)
}, 1000)

setTimeout(() => {
  console.info(`Start timeout date time at: ${new Date()}`)
}, 1000)

let iteration = 1
const intervalFunc = () => {
  console.info(`Iteration ${iteration}, at: ${new Date()}`)
  iteration += 1
  if (iteration % 5 == 0) {
    console.info(`-----Break------`)
  }
}
setInterval(intervalFunc, 2000)