import cluster from "cluster"
import os from "os"
import http from "http"

if (cluster.isPrimary) {
  console.info(`Run cluster primary in ${process.pid}`)
  
  for (let i=0; i < os.cpus().length; i++) {
    cluster.fork()
    cluster.addListener("exit", (worker, code, signal) => {
      cluster.fork()
    })
  }
}

if (cluster.isWorker) {
  console.info(`Run cluster worker in ${process.pid}`)
  const server = http.createServer((request, response) => {
    response.write(`Welcome to my API, pid: ${process.pid}`)
    response.end()
    process.exit()
  })
  server.listen("3000")
}