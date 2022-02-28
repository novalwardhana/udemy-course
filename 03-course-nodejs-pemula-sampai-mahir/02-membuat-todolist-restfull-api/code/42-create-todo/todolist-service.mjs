class TodoListService {
  todoList = ["Ngoding", "Dolan"]

  jsonGetTodoList() {
    return JSON.stringify({
      code: 200,
      status: "OK",
      data: this.todoList.map((value, index) => {
        return {
          id: index,
          todo: value
        }
      })
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
}

export {TodoListService}