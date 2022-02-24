import readline from "readline";
import fs from "fs"
import { Console } from "console";

const fsLog = fs.createWriteStream("application-log3.log")
const log = new Console({
  stdout: fsLog,
  stderr: fsLog
})
const input = readline.createInterface({
  input: process.stdin,
  output: process.stdout
})
input.question("Who's your name?", (name) => {
  log.info(name)
  input.question("age?", (age) => {
    log.info(age)
    input.question("address?", (address) => {
      log.info(address)
      input.question("Favorite club?", (club) => {
        log.info(club)
        input.question("status?", (status) => {
          log.info(status)
          input.close()
        })
      })
    })
  })
})