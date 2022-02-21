/*
  Events
  a. Events adalah standard library di nodejs yang bisa digunakan sebagai implementasi event listener
  b. Di dalam events terdapat event EventEmitter yang bisa digunakan untuk menampung data listener per jenis events
  c. Kita bisa juga melakukan emmit untuk mentrigger jenis event dan mengirim data ke event tersebut
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/events.html
*/

import {EventEmitter} from "events"
const emitter = new EventEmitter()

emitter.addListener("hello", (name) => {
  console.info(`Hello ${name}`)
})
emitter.addListener("hello", (name) => {
  console.info(`Holla ${name}`)
})
emitter.emit("hello", "Noval wardhana")

emitter.addListener("getMax", (...datas) => {
  let result = datas[0]
  for (const data of datas) {
    if (data > result) {
      result = data
    }
  }
  console.info(`maxData: ${result}`)
})
emitter.addListener("getMin", (...datas) => {
  let result = datas[0]
  for (const data of datas) {
    if (data < result) {
      result = data
    }
  }
  console.info(`minData: ${result}`)
})
emitter.addListener("getSum", (...datas) => {
  let result = 0
  for (const data of datas) {
    result += data
  }
  console.info(`getSum: ${result}`)
})
emitter.emit("getMax", 7, 8, 9, 10, 6, 9, 8, 6, 9, 9)
emitter.emit("getMin", 7, 8, 9, 10, 6, 9, 8, 6, 9, 9)
emitter.emit("getSum", 7, 8, 9, 10, 6, 9, 8, 6, 9, 9)