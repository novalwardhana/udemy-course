const getMessage = (message) => {
  for (let i = 0; i < message.data; i++) {
    const message = `Worker5: ${i}`
    postMessage(message)
  }
}
addEventListener("message", getMessage)