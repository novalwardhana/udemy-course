'use strict'

class Club {

  constructor(name, state) {
    this.name = name
    this.state = state
  }

  get Name() {
    return this.name
  }

  get State() {
    return this.state
  }

  set Name(value) {
    this.name = value
  }

  set State(value) {
    this.state = value
  }

}