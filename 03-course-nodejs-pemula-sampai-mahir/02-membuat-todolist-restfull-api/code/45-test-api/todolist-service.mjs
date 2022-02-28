export class TodoListService {

  todoList = ["Dolan", "Mancing"]

  jsonGetTodoList() {
    return JSON.stringify({
      code: 200,
      message: "OK",
      data: this.todoList
    })
  }

  getTodoList(request, response) {
    response.write(this.jsonGetTodoList())
    response.end()
  }

  createTodoList(request, response) {
    request.addListener("data", (data) => {
      const body = JSON.parse(data.toString())
      this.todoList.push(body.todo)
      response.write(this.jsonGetTodoList())
      response.end()
    })
  }

  updateTodoList(request, response) {
    request.addListener("data", (data) => {
      const body = JSON.parse(data.toString())
      if (this.todoList[body.id]) {
        this.todoList[body.id] = body.todo
      }
      response.write(this.jsonGetTodoList())
      response.end()
    })
  }

  deleteTodoList(request, response) {
    request.addListener("data", (data) => {
      const body = JSON.parse(data.toString())
      if (this.todoList[body.id]) {
        this.todoList.splice(body.id, 1)
      }
      response.write(this.jsonGetTodoList())
      response.end()
    })
  }

}