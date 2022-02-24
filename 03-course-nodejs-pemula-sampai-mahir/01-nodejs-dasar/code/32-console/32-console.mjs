/*
  Console
  a. Secara global, object console bisa kita gunakan tanpa harus import module. Dan console melakukan printe text nya ke stdout
  b. Namun kita bisa membuat object console sendiri jika kita mau
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/console.html

  Memasukkan console ke dalam file
*/

import {Console} from "console"
import fs from "fs"

const fileLog = fs.createWriteStream("application-log.log")
const log = new Console({
  stdout: fileLog,
  stderr: fileLog
})

const table = [
  {
    no: 1,
    club: "Manchester United",
    coach: "Ragnick"
  },
  {
    no: 2,
    club: "Chelsea",
    coach: "Tuchel"
  },
  {
    no: 3,
    club: "Liverpool",
    coach: "Juergen Klopp"
  }
]
log.info("Noval")
log.info("Wardhana")
log.table(table)

const fileLog2 = fs.createWriteStream("application-log2.log")
const log2 = new Console({
  stdout: fileLog2,
  stderr: fileLog2
})
log2.info("Hello World")
log2.info("Good Morning")
for (let i = 1; i <= 10; i++) {
  let result = ""
  for (let j = 1; j<= i; j++) {
    result += "*"
  }
  log2.info(result)
}