import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker2 in therad-${threadId}`)
  
  let result = message[0]
  for (const data of message) {
    if (data > result) {
      result = data
    }
  }
  parentPort.postMessage(result)

  parentPort.close()
})