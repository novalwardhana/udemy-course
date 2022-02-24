import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker9 in thread-${threadId}`)

  const result = message.map(element => Math.sqrt(element))
  parentPort.postMessage(result)
  
  parentPort.close()
})