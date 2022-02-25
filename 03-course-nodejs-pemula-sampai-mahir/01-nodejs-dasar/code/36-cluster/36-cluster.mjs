/*
  Cluster
  a. NodeJS berjalan di single thread, kecuali jika kita membuat thread manual menggunakan worker thread, tapi tetap dalam satu proses
  b. NodeJS memiliki standard library bernama cluster, dimana kita bisa menjalankan beberapa process NodeJS secara sekaligus
  c. Cocok digunakan ketika kita menggunakan CPU yang multicore, sehingga semua core bisa kita utilisasi dengan baik, misal kita jalankan process NodeJS
    sejumlah CPU core
  d. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/cluster.html

  Cluster Primary dan Worker
  a. Di dalam cluster, terdapat 2 jenis aplikasi yaitu Primary dan Worker
  b. Primary biasanya digunakan sebagai koordinator atau manajer untuk para Worker
  c. Sedangkan Worker sendiri adalah aplikasi yang menjalankan tugasnya
*/

import cluster from "cluster"
import process from "process"
import os from "os"
import http from "http"

if (cluster.isMaster) {
  console.info(`Cluster primary, pid: ${process.pid}`)
  for (let i = 0; i < os.cpus().length; i++) {
    cluster.fork()
  }

  cluster.addListener("exit", (worker) => {
    console.info(`Worker exit: ${worker.id}`)
    cluster.fork()
  })
}

if (cluster.isWorker) {
  console.info(`Cluster Worker, pid: ${process.pid}`)
  const server = http.createServer((request, response) => {
    response.write(`Response from process: ${process.pid}`)
    response.end()
    process.exit()
  })
  server.listen("3000")
}