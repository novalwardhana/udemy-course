class myClass {

  constructor(name) {
    this.name = name
  }

  factorial(number) {
    let result = 0
    if (number == 1) {
      return 1
    } else {
      result += number * this.factorial(number - 1)
    }
    return result
  }
}

export default myClass