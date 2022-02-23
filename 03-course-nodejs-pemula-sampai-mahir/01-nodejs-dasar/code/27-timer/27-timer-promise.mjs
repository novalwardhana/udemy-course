import timers from "timers/promises"

console.info(`Date now at: ${new Date()}`)
const result1 = await timers.setTimeout(5000, "Noval")
console.info(`Date now at: ${new Date()}`)
console.info(`result: ${result1}`)

console.info(`Date now at: ${new Date()}`)
const result2 = await timers.setTimeout(2000, "Tottenham")
console.info(`Date now at: ${new Date()}`)
console.info(`result2: ${result2}`)

console.info(`Date now at: ${new Date()}`)
const result3 = await timers.setTimeout(3000, "NodeJS")
console.info(`Date now at: ${new Date()}`)
console.info(`result3: ${result3}`)

let maxIteration = 10
let iteration = 0
for await (const result of timers.setInterval(1000, "ignored")) {
  console.info(`Iteration ${iteration}, at: ${new Date()}`)
  if (maxIteration == iteration) {
    break
  }
  iteration++
}

let maxLoop = 8
let loop = 0
for await (const result of timers.setInterval(1500, "ignored value")) {
  console.info(`loop ${loop}, at: ${new Date()}`)
  if (loop == maxLoop) {
    break
  }
  loop++
}