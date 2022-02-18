const getMessage = (message) => {
  let i = 0
  while (i < message.data) {
    const message = `worker4: ${i}`
    postMessage(message)
    i++
  }
}
addEventListener("message", getMessage)