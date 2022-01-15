export class Person {

  constructor(name) {
    this.name = name
  }

  static sayHai(guest) {
    console.info(`Haii ${guest}`)
  }

  sayGoodBye = function(guest) {
    console.info(`Goodbye ${guest}`)
  }

  sayHello(guest) {
    console.info(`Hello ${guest}, my name is ${this.name}`)
  }
}