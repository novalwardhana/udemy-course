import {threadId, parentPort} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker3 in thread-${threadId}`)

  let result = message[0]
  for (const data of message) {
    if (data < result) {
      result = data
    }
  }
  parentPort.postMessage(result)

  parentPort.close()
})