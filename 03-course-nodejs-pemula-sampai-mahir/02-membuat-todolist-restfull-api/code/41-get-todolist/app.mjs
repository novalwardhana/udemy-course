import http from "http"
import { TodoListService } from "./todolist-service.mjs"

const service = new TodoListService()
const server = http.createServer((request, response) => {

  response.setHeader("Content-Type", "application/json")

  if (request.method === "GET") {
    service.getTodoList(request,  response)
  }
})
server.listen("3000")