import http, { get } from "http"

const server = http.createServer((request, response) => {

  if (request.method === "POST" && request.url === "/statistic") {
    request.addListener("data", (data) => {
      const params = data.toString()
      const getData = statistic(JSON.parse(params))
      getData.then((respData) => {
        response.write(Buffer.from(JSON.stringify(respData)))
        response.end()
      }).catch(error => {
        response.write("An error occured")
        response.end()
      })
    })
  } else if (request.method === "POST" && request.url === "/square") {
    request.addListener("data", (data) => {
      const params = data.toString()
      const arrParams = JSON.parse(params)
      const result = arrParams.map(element => Math.pow(element, 2))
      response.write(Buffer.from(JSON.stringify(result)))
      response.end()
    })
  } else {
    response.write("Welcome to my API")
    response.end()
  }
  
})
server.listen("3000")

const maxData = async (datas) => {
  let max = datas[0]
  for (const data of datas) {
    if (data > max) {
      max = data
    }
  }
  return max
}
const minData = async (datas) => {
  let min = datas[0]
  for (const data of datas) {
    if (data < min) {
      min = data
    }
  }
  return min
}
const sumData = async (datas) => {
  let sum = 0
  for (const data of datas) {
    sum += data
  }
  return sum
}
const statistic = async (datas) => {
  const max = await maxData(datas)
  const min = await minData(datas)
  const sum = await sumData(datas)
  const mean = sum / datas.length
  const result = {
    "max": max,
    "min": min,
    "sum": sum,
    "mean": mean
  }
  return result
}