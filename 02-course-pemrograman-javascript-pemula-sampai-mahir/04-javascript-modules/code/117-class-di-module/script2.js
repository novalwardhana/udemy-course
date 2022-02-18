export class Statistic {

  constructor(...datas) {
    this.datas = datas
  }

  sum() {
    let result = 0
    for (const data of this.datas) {
      result += data
    }
    return result
  }

  maxNumber() {
    let result = 0
    for (const data of this.datas) {
      if (result < data) {
        result = data
      }
    }
    return result
  }

  mean() {
    const sumDatas = this.sum()
    const length = this.datas.length
    const result = sumDatas / length
    return result
  }

}