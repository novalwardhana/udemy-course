/*
  Report
  a. Report merupakan fitur yang terdapat di NodeJS untuk membuat laporan secara otomatis dalam file
    ketika terjadi sesuatu pada aplikasi NodeJS kita
  b. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/report.html 
*/

import process from "process"

process.report.reportOnFatalError = true
process.report.reportOnSignal = true
process.report.reportOnUncaughtException = true
process.report.filename = "report.json"

function sampleError() {
  throw new Error("Terjadi error")
}

const getMaxNumber = (...datas) => {
  try {
    if (datas.length == 0) {
      throw new Error("Data is null")
    }

    let result = datas[0]
    for (const data of datas) {
      if (data > result) {
        result = data
      }
    }
    return result
  } catch (error) {
    return {errorMessage: error.message}
  }
}
const maxNumber = getMaxNumber()
if (maxNumber.errorMessage !== undefined) {
  throw new Error(maxNumber.errorMessage)
} else {
  console.info(`maxNumber: `, maxNumber)
}