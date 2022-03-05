/*
  Package:
  npm install --save express -> install framework express js
  npm install --save-dev nodemon -> refresh every change
  npm install --save mysql -> handle mysql connection
  npm install --save body-parser -> handle json request body
  npm install --save swagger-jsdoc swagger-ui-express

  package.json
  package.josn -> script -> start: nodemon index.js

*/

const express = require("express")
const app = express()
const bodyParser = require("body-parser")
const usersRoutes = require("./routes/users.route")
const postsRoutes = require("./routes/posts.route")
const swaggerJsdoc = require("swagger-jsdoc")
const swaggerUi = require("swagger-ui-express")

app.use(bodyParser.json())

/* Swagger documentation */
const swaggerOption = {
  swaggerDefinition:(swaggerJsdoc.Options = {
    info: {
      title: "my-posts",
      description: "API Documentation",
      contact: {
        name: "Developer"
      },
      servers: ["http://localhost:3000/"]
    }
  }),
  apis: ["index.js", "./routes/*.js"]
}
const swaggerDocs = swaggerJsdoc(swaggerOption)
app.use("/api-docs", swaggerUi.serve, swaggerUi.setup(swaggerDocs))
/* End swagger documentation */

app.use("/test", (req, res) => {
  console.info(`Receiving request`)
  res.status(200).send("Success 2")
})
app.use("/users", usersRoutes)
app.use("/posts", postsRoutes)

app.listen(3000, () => {
  console.info("Start nodejs application")
})