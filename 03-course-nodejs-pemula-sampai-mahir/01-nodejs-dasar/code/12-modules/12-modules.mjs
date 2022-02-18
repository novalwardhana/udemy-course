/*
  Modules
  a. Standard library yang terdapat di NodeJS bisa kita gunakan seperti layaknya Javascript Modules
  b. Karena NodeJS menggunakan Modules, jika kita ingin menggunakan Modules, kita juga perlu
    memberi tahu bahwa file Javascript kita menggunakan Modules,
    caranya dengan mengubah nama file dari .js menjadi .mjs
*/

import os from "os"
import myClass from "./example-module.mjs"
import MyMath from "./example-module2.mjs"

console.info(`OS platforms: `, os.platform())
console.table(os.cpus())
console.info(`OS Hostname: `, os.hostname())
console.info(`OS Homedir: `, os.homedir())

const example = new myClass("noval")
console.info(example.name)
console.info(example.factorial(10))

const myMath = new MyMath(6, 5, 8, 7, 9, 8, 9, 4, 7, 9)
const myMathDatas = myMath.datas
const myMathMax = myMath.maxNumber()
const myMathMin = myMath.minNumber()
const myMathTotal = myMath.sumValue()
const myMathMean = myMath.meanValue()
console.info("\nMath Function")
console.info(`Data: `, myMathDatas)
console.info(`Max Value: ${myMathMax}`)
console.info(`Min Value: ${myMathMin}`)
console.info(`Sum Value: ${myMathTotal}`)
console.info(`Mean Value: ${myMathMean}`)