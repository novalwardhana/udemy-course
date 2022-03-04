const { createPool } = require("mysql")

const db = createPool({
  host: "localhost",
  port: "3306",
  user: "root",
  password: "noval",
  database: "my-posts",
  connectionLimit: 10
})

module.exports = db