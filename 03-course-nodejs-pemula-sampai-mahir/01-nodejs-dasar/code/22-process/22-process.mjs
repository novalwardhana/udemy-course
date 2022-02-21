/*
  Process
  a. Process adalah standard library yang digunakan untuk mendapatkan informasi proses NodeJS yang sedang berjalan
  b. Process merupakan instance dari event emitte, sehingga kita bisa menambahkan listener ke dalam process
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/process.html
*/

import process from "process"

process.addListener("exit", (number) => {
  console.info(`Exit with code ${number}`)
})
process.addListener("beforeExit", (number, date) => {
  console.info(`Before exit ${date}, and number ${number}`)
})

console.info(`version: `, process.version)
console.info(`argv: `,process.argv)
console.table(process.env)
console.table(process.report)
console.info(`cpu Usage: `, process.cpuUsage())

process.exit(21)

console.info(`Test call process`)