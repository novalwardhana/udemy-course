import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker10 in thread-${threadId}`)

  const result = message.map(element => Math.pow(element, 3))
  parentPort.postMessage(result)

  parentPort.close()
})