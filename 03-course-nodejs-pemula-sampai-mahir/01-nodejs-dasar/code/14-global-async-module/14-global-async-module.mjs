/*
  GLobal Async Module
  a. Saat kita belajar Javascript async-await, biasanya kita perlu membuat terlebih dahulu function
    yang kita tandai sebagai async
  b. Saat kita menggunakan module, secara default global code adalah async, oleh karena itu kita bisa
    menggunakan async-await
  c. Kecuali jika kita membuat function, maka function tersebut harus kita tandai sebagai async jika 
    ingin menggunakan async-await
  d. Hanya bisa digunakan pada file dengan ekstensi .mjs dan nodeJS versi 16 keatas
*/

async function samplePromise() {
  return Promise.resolve(`Hello world`)
}
const name = await samplePromise()
console.info(name)

async function sayMyName() {
  return "Haloo Novalita"
}
const myName = await sayMyName()
console.info(myName)

const mathSquare = async (...datas) => {
  let result = datas.map(element => Math.pow(element, 2))
  return result
}
const squareData = await mathSquare(10, 2, 5, 4, 8, 7)
console.info(`squareData: `, squareData)

const getSumData = async (datas) => {
  let result = 0
  for (const data of datas) {
    result += data
  }
  return result
}
const getMeanData = async (datas) => {
  const sum = await getSumData(datas)
  const count = datas.length
  const result = sum / count
  return result
}
const getMaxData = async (datas) => {
  let result = 0
  for (const data of datas) {
    if (result < data) {
      result = data
    }
  }
  return result
}
const getMinData = async (datas) => {
  let result = datas[0]
  for (const data of datas) {
    if (result > data) {
      result = data
    }
  }
  return result
}
const getStatisticData = async (...datas) => {
  const sum = await getSumData(datas)
  const mean = await getMeanData(datas)
  const max = await getMaxData(datas)
  const min = await getMinData(datas)
  const result = {
    "sum": sum,
    "mean": mean,
    "max": max,
    "min": min
  }
  return result
}
const statisticData = await getStatisticData(9, 5, 3, 8, 6, 9, 9, 7, 8, 9)
console.info(`statisticData: `, statisticData)
const statisticData2 = await getStatisticData(3, 8, 9, 4, 2, 7)
console.info(`statisticData2: `, statisticData2)
const statisticData3 = await getStatisticData(8, 4, 6, 19, 3, 4)
console.info(`statisticData3: `, statisticData3)