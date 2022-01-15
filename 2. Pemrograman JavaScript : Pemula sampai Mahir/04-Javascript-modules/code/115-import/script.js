export function sayHello(name) {
  console.info(`Hello ${name}`)
}

export function sayGoodBye(name) {
  console.info(`Goodbye ${name}`)
}

export function maxNumber(...datas) {
  let maxNumber = 0
  for (const data of datas) {
    if (maxNumber < data) {
      maxNumber = data
    }
  }
  return maxNumber
}