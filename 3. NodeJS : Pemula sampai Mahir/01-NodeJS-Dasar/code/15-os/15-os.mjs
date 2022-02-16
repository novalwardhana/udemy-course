/*
  OS
  Adalah standard library yang bisa digunakan untuk memperoleh informasi tentang sistem operasi yang digunakan
  Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/os.html
*/

import os, { freemem } from "os"

const platform = os.platform()
const arch = os.arch()
const cpus = os.cpus()
const uptime = os.uptime()
const totalMemory = os.totalmem()
const freeMemory = os.freemem()
const userInfo = os.userInfo()
const network = os.networkInterfaces()

console.info(`Platform: `, platform)
console.info(`Arch: `, arch)
console.table(cpus)
console.info(`Uptime: ${uptime}`)
console.info(`Total memory: ${totalMemory}`)
console.info(`Free memory: ${freeMemory}`)
console.info(`userInfo: `, userInfo)
console.info(`network: `, network)