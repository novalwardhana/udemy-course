const firstName = "Lionel"
const lastName = "Messi"

function sayHello(guest) {
  console.info(`Hello ${guest}`)
}

const sumData = (...datas) => {
  let result = 0
  for(const data of datas) {
    result += data
  }
  return result
}

class Company {
  name = "Stark Company"
}

export {firstName, lastName, sayHello, sumData, Company}