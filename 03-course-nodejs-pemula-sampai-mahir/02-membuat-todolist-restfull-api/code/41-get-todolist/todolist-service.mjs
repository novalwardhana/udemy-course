class TodoListService {

  todoList = ["Ngoding", "Touring", "Eating"]

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

}

export {TodoListService}