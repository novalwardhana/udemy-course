/*
  Require
  a. NodeJS awalnya tidak menggunakan Javascript Modules (belum rilis), namun sekarang NodeJS sudah bisa menggunakan Javascript Modules
    dan sangat direkomendasikan menggunakannya
  b. Sebelum menggunakan Modules, NodeJS menggunakan require() untuk melakukan import file
  c. 
*/

const os = require("os")
const MyMath = require("./example-require.js")

console.info(`OS platform: `, os.platform())
console.table(os.cpus())

const myMath = new MyMath(6, 5, 8, 7, 9, 8, 9, 4, 7, 9)
const myMathSquare = MyMath.square(6, 5, 8, 7, 9, 8, 9, 4, 7, 9)
const myMathSquareRoot = MyMath.squareRoot(6, 5, 8, 7, 9, 8, 9, 4, 7, 9)
const myMathMean = myMath.meanValue()
const myMathVariance = myMath.varianceValue()
console.info(`myMath: `, myMath)
console.info(`myMathSquare: `, myMathSquare)
console.info(`myMathSquareRoot: `, myMathSquareRoot)
console.info(`myMathMean: ${myMathMean}`)
console.info(`mYMathVariance: ${myMathVariance}`)