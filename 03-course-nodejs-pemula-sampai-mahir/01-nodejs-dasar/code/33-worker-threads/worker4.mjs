import {threadId, parentPort} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker4 in therad-${threadId}`)

  let result = 0
  for (const data of message) {
    result += data
  }
  parentPort.postMessage(result)

  parentPort.close()
})