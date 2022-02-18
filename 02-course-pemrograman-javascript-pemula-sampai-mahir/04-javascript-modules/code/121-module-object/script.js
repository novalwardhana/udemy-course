export const firstName = "Novalita"
export const middleName = "Kusuma"
export const lastName = "Wardhana"

export const sayHello = (name) => {
  console.info(`Hellow ${name}`)
}

export class Person {
  constructor(name) {
    this.name = name
  } 
}