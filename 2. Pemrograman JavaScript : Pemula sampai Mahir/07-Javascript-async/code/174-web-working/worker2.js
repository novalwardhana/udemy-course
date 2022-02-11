addEventListener("message", (message) => {
  const total = message.data
  for (let i = 0; i < total; i++) {
    const message = `worker2: ${i}`
    postMessage(message)
  }
})