import {threadId, Worker} from "worker_threads"

const worker2 = new Worker("./worker2.mjs")
const worker3 = new Worker("./worker3.mjs")
const worker4 = new Worker("./worker4.mjs")
const worker5 = new Worker("./worker5.mjs")
const worker6 = new Worker("./worker6.mjs")
const worker7 = new Worker("./worker7.mjs")
const worker8 = new Worker("./worker8.mjs")
const worker9 = new Worker("./worker9.mjs")
const worker10 = new Worker("./worker10.mjs")

worker2.addListener("message", (message) => {
  console.info(`max: ${message}, execution by thread-${threadId}`)
})
worker3.addListener("message", (message) => {
  console.info(`min: ${message}, execution by thread-${threadId}`)
})
worker4.addListener("message", (message) => {
  console.info(`sum: ${message}, execution by thread-${threadId}`)
})
worker5.addListener("message", (message) => {
  console.info(`mean: ${message}, execution by thread-${threadId}`)
})
worker6.addListener("message", (message) => {
  console.info(`even numbers: ${message}, execution by thread-${threadId}`)
})
worker7.addListener("message", (message) => {
  console.info(`odd number: ${message}, execution by thread-${threadId}`)
})
worker8.addListener("message", (message) => {
  console.info(`Square number: ${message}, execution by thread-${threadId}`)
})
worker9.addListener("message", (message) => {
  console.info(`Square root number: ${message}, execution by therad-${threadId}`)
})
worker10.addListener("message", (message) => {
  console.info(`Cubic: ${message}, execution by thread=${threadId}`)
})

const datas = [8, 5, 4, 9, 2, 5, 4, 9, 4, 5, 8, 7]
worker2.postMessage(datas)
worker3.postMessage(datas)
worker4.postMessage(datas)
worker5.postMessage(datas)
worker6.postMessage(datas)
worker7.postMessage(datas)
worker8.postMessage(datas)
worker9.postMessage(datas)
worker10.postMessage(datas)