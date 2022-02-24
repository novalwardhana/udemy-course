/*
  Util
  a. Merupakan standard library di NodeJS yang berisikan utility2 yang bisa kita gunakan untuk mempermudah koide nodejs
  b. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/util.html
*/

import util from "util"

const firstName = "Noval"
const lastName = "Wardhana"
console.info(`Hello ${firstName} ${lastName}`)
console.info(util.format("hello %s %s", firstName, lastName)) // Digunakan sebelum adanya fitur backtick

const firstClub = "Manchester United"
const secondClub = "Juventus"
const merge = util.format("%s and %s", firstClub, secondClub)
console.info(merge)

const person = {
  firstName: "Jadon",
  lastName: "Sancho"
}
console.info(`Person ${JSON.stringify(person)}`)
console.info(util.format("Person %j", person))

const club = {
  name: "Lazio",
  city: "Roma",
  country: "Italy"
}
console.info(util.format("club %j", club))

const validation = util.types.isStringObject(person)
const validation2 = util.types.isDate(new Date())
console.info(validation)
console.info(validation2)

const data1 = {
  name: "Noval",
  age: 25
}
const data2 = {
  name: "Noval",
  age: 25
}
const data3 = {
  age: 25,
  name: "Noval"
}
const data4 = {
  name: "Noval"
}
const validation3 = util.isDeepStrictEqual(data1, data2)
console.info(validation3)
const validation4 = util.isDeepStrictEqual(data1, data3)
console.info(validation4)
const validation5 = util.isDeepStrictEqual(data2, data4)
console.info(validation5)