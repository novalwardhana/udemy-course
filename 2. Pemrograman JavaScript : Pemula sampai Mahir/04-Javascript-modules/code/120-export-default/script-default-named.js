const firstName = "Manchester"

const lastName = "United"

const sayHello = (guest) => {
  console.info(`Hello bro, ${guest}`)
}

class Person {
  constructor(name, age, address) {
    this.name = name
    this.age = age
    this.address = address
  }
}

export default Person
export {firstName, lastName, sayHello}