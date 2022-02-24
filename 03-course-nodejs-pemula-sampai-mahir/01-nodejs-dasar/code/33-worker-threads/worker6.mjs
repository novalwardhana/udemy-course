import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {
  console.info(`worker6 in thread-${threadId}`)

  let result = message.filter(element => element % 2 == 0)
  parentPort.postMessage(result)

  parentPort.close()
})