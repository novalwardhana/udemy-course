const name = "Novalita Kusuma Wardhana"

function sayHello(guest) {
  console.info(`Hello ${guest}`)
}

const sayGoodBye = function(guest) {
  console.info(`Goodbye ${guest}`)
}

class statistic {

  constructor(...datas) {
    this.datas = datas
  }

  sumData() {
    let result = 0
    for (const data of this.datas) {
      result += data
    }
    return result
  }

  meanData() {
    const sum = this.sumData()
    const length = this.datas.length
    const result = sum / length
    return result
  }

  varianceData() {
    const length = this.datas.length
    const mean = this.meanData()
    let sumDataMean = 0
    for (const data of this.datas) {
      sumDataMean += Math.pow((data - mean), 2)
    }
    const result = sumDataMean / length
    return result
  }

  stdDeviationData() {
    const result = Math.sqrt(this.varianceData())
    return result
  }
}

export {name, sayHello, sayGoodBye, statistic}