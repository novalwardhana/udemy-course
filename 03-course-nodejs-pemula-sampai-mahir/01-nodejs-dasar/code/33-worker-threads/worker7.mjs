import {parentPort, threadId} from "worker_threads"

parentPort.addListener("message", (message) => {

  console.info(`worker7 in therad-${threadId}`)

  const result = message.filter(element => element % 2 == 1 )
  parentPort.postMessage(result)

  parentPort.close()
})