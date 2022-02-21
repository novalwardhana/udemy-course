/*
  GLobals
  a. Di dalam nodejs terdapat library berupa variable atau function yang secara global bisa diakses dimana saja,
    tanpa harus melakukan import
  b. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/globals.html
*/

setTimeout(() => {
  console.info(`Hello globals`)
}, 2000)

setInterval(() => {
  const newDate = new Date().toString()
  console.info(`Set intervals: ${newDate}`)
}, 2000)