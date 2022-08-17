const express = require("express")
const app = express()

app.use(express.json())

/* Setting up confih.env file variables */
const dotenv = require("dotenv")
dotenv.config({path: './config/config.env'})

/* Hadnle uncaught exception */
process.on("uncaughtException", err => {
  console.info(`Error: ${err.message}`)
  console.info("Handle error uncaught exception")
  process.exit(1)
})
//console.log(abc)

/* Database handler */
const connectDatabase = require("./config/database")
connectDatabase()

/* Error middleware */
const errorMiddleware = require("./middlewares/error")

/* Middleware */
const middleware = (req, res, next) => {
  console.info("Use middleware")
  req.exampleUser = "Noval Wardhana"
  req.requestMethod = req.method
  next()
}
app.use(middleware)

/* Jobs Handler */
const jobs = require("./routes/jobs")
app.use("/api/v1", jobs)

const ErrorHandler = require("./utils/errorHandler")
/* Handle unhandled route */
app.all("*", (req, res, next) => {
  next(new ErrorHandler(`${req.originalUrl}: not found`, 404))
})

/* Error middleware: diletakkan di akhir */
app.use(errorMiddleware)

const port = process.env.PORT
console.info(port)
const server = app.listen(port, () => {
  console.info(`Start nodejs server, in PORT: ${process.env.PORT}, and NODE_ENV: ${process.env.NODE_ENV}`)
})

/* Handle Unhandled promise rejection when mongo env not match */
process.on("unhandledRejection", err => {
  console.info(`Error: ${err.message}`)
  console.info("Shutting down the server due to unhandled promise rejection")
  server.close(() => {
    process.exit(1)
  })
})