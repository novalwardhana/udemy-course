const events = require("events")

let eventEmitter = new events.EventEmitter()
eventEmitter.on("connection", () => {
  console.info(`Connection event driven`)
})
eventEmitter.emit("connection")

eventEmitter.on("callme", (data) => {
  console.info(`callme event driven: `, data)
})
eventEmitter.emit("callme", [1, 4, 9, 16, 25])