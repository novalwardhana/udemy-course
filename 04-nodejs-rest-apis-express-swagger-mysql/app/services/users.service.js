
const db = require("../config/db.config")

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