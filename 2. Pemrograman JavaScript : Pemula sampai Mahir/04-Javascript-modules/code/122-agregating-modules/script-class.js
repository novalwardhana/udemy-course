class Statsictic {
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
    const len = this.datas.length
    const result = sum / len
    return result
  }
}

class Probability {

  permutation(number) {
    let result = 0
    if (number <= 1) {
      result += 1
    } else {
      result += number * this.permutation(number - 1)
    }
    return result
  }

  combination(from, to) {
    const fromPermutation = this.permutation(from)
    const toPermutation = this.permutation(to)
    const diff = this.permutation(from - to)
    const result = fromPermutation / (toPermutation * diff)
    return result
  }

}

export {Statsictic, Probability}