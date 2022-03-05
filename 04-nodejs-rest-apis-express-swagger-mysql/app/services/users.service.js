
const db = require("../config/db.config")

module.exports.login = (data, callback) => {
  db.query("select * from users where emailId = ? and password = ?", [data.emailId, data.password], (error, results, fields) => {
    if (error) {
      return callback(error)
    }
    if (results.length === 0) {
      return callback(new Error("Invalid credentials"))
    }
    return callback(null, results)
  })
}

module.exports.register = (data, callback) => {
  db.query("insert into users (firstName, lastName, emailId, password) values (?, ?, ?, ?) ", [data.firstName, data.lastName, data.emailId, data.password],
    (error, results, fields) => {
      if (error) {
        return callback(error, null)
      }
      return callback(null, results)
    }
  )
}