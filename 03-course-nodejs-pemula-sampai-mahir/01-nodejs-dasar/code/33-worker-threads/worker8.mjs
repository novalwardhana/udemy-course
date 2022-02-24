import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker8 in theread-${threadId}`)

  const result = message.map(element => Math.pow(element, 2))
  parentPort.postMessage(result)

  parentPort.close()
})