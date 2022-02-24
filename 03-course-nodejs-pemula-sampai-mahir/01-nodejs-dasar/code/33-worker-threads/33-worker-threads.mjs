/*
  Worker Threads
  a. Worker Threads adalah standard library yang bisa kita gunakan untuk menggunakan thread ketika mengeksekusi Javascript secara parallel
  b. Worker Threads sangat cocok ketika kita membuat kode program yang butuh jalan secara parallel,
    dan biasanya kasusnya adalah ketika kode program kita membutuhkan proses CPU intensive,
    seperti misalnya enkripsi atau kompresi
  c. Cara kerja Worker Threds mirip dengan Worker di Javascript Web API
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/worker_threads.html
*/

import {threadId, Worker} from "worker_threads"

const worker1 = new Worker("./worker.mjs")
const worker2 = new Worker("./worker.mjs")

worker1.addListener("message", (message) => {
  console.info(`Thread ${threadId} worker1 receiving message ${message}`)
})
worker2.addListener("message", (message) => {
  console.info(`Thread ${threadId} worker2 receiving message ${message}`)
})

worker1.postMessage(10)
worker2.postMessage(10)