class MyMath {

  constructor(...datas) {
    this.datas = datas
  }

  static square(...datas) {
    const result = datas.map(element => Math.pow(element, 2))
    return result
  }

  static squareRoot(...datas) {
    const result = datas.map(element => Math.sqrt(element))
    return result
  }

  sumValue() {
    let sum = 0
    for (const data of this.datas) {
      sum += data
    }
    return sum
  }

  meanValue() {
    const sum = this.sumValue()
    const count = this.datas.length
    const result = sum / count
    return result
  }

  varianceValue() {
    const mean = this.meanValue()
    let total = 0
    for (const data of this.datas) {
      total += Math.pow((data - mean), 2)
    }
    const result = total / this.datas.length
    return result
  }

}

module.exports = MyMath