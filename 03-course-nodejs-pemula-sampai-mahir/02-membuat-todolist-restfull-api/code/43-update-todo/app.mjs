import http from "http"
import { Todolist } from "./todolist-service.mjs"

const service = new Todolist()
const server = http.createServer((request, response) => {
  response.setHeader("Content-Type", "application/json")

  if (request.method === "GET") {
    service.getTodoList(request, response)
  } else if (request.method === "POST") {
    service.creatTodoList(request, response)
  } else if (request.method === "PUT") {
    service.updateTodoList(request, response)
  } 
})
server.listen("3000")