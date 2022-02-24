import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker5 in therad-${threadId}`)

  let sum = 0
  for (const data of message) {
    sum += data
  }
  const result = sum / message.length
  parentPort.postMessage(result)

  parentPort.close()
})