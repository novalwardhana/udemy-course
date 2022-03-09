const fs = require("fs")

/* Synchronous */
const fileSync = fs.readFileSync("02-callback.txt")
console.info(`fileSync: `, fileSync.toString())
console.info(`Program ended synchronous`)

/* Asynchronous */
fs.readFile("02-callback.txt", (error, data) => {
  console.info(`fileAsync: `, data.toString())
})
console.info(`Program ended asynchronous`)