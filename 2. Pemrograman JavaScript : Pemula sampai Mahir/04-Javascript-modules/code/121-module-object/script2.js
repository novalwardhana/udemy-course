export class Club {

  constructor(name, league) {
    this.name = name
    this.league = league
  }

  get thisName() {
    return this.name
  }

  set thisName(value) {
    this.name = value
  }

  get thisLeague() {
    return this.league
  }

  set thisLeague(value) {
    this.league = this.league
  }

}

export class Probability {

  permutation(number) {
    if (number <= 1) {
      return 1
    } else {
      return number * this.permutation(number -1)
    }
  }

  combination(from, to) {
    const fromPermutation = this.permutation(from)
    const toPermutation = this.permutation(to)
    const differencePermutation = this.permutation(from - to)
    const result = fromPermutation / (toPermutation * differencePermutation)
    return result
  }

}