class Todolist {

  todoList = ["Touring", "Jogging"]

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

  creatTodoList(request, response) {
    request.addListener("data", (data) => {
      const params = JSON.parse(data.toString())
      this.todoList.push(params.todo)
      response.write(this.jsonGetTodoList())
      response.end()
    })
  }

  updateTodoList(request, response) {
    request.addListener("data", (data) => {
      const body = JSON.parse(data.toString())
      this.todoList[body.id] = body.todo
      response.write(this.jsonGetTodoList())
      response.end()
    })
  }

}

export { Todolist }