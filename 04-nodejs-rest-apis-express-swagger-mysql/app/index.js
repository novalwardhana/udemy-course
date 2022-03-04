/*
  Package:
  npm install --save express -> install framework express js
  npm install --save-dev nodemon -> refresh every change
  npm install --save mysql -> handle mysql connection
  npm install --save body-parser -> handle json request body

  package.json
  package.josn -> script -> start: nodemon index.js
*/

const express = require("express")
const app = express()
const bodyParser = require("body-parser")
const usersRoutes = require("./routes/users.route")

app.use(bodyParser.json())

app.use("/test", (req, res) => {
  console.info(`Receiving request`)
  res.status(200).send("Success 2")
})
app.use("/users", usersRoutes)

app.listen(3000, () => {
  console.info("Start nodejs application")
})