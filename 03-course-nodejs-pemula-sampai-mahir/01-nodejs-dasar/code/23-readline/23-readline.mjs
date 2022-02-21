/*
  Readline
  a. Standard library yang digunakan untuk membaca input
  b. Di versi NodeJS 16 hanya mendukung callback, sedangkan di versi NodeJS 17 sudah mendukung callback
  c. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/readline.html
*/

import readline from "readline"
import process from "process"

const input = readline.createInterface({
  input: process.stdin,
  output: process.stdout
})

input.question("Who's your name?", (name) => {
  console.info(`Name: ${name}`)
  input.question("Your favorite food?", (name) => {
    console.info(`Food: ${name}`)
    input.question("Your hometown?", (name) => {
      console.info(`Hometown: ${name}`)
      input.question("Your age?", (age) => {
        console.info(`Age: ${age}`)
        input.close()
      })
    })
  })
})