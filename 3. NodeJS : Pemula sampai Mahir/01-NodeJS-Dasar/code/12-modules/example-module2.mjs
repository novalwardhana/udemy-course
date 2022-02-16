class MyMath {

  constructor(...datas) {
    this.datas = datas
  }

  maxNumber() {
    let max = 0
    for (const data of this.datas) {
      if (data > max) {
        max = data
      }
    }
    return max
  }

  minNumber() {
    let min = this.datas[0]
    for (const data of this.datas) {
      if (data < min) {
        min = data
      }
    }
    return min
  }

  sumValue() {
    let sum = 0
    for (const data of this.datas) {
      sum += data
    }
    return sum
  }

  meanValue() {
    let sum = this.sumValue()
    let count = this.datas.length
    let mean = sum / count
    return mean
  }

}

export default MyMath