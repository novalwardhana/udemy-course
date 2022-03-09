const express = require("express")
const app = express()

app.use(express.json())

app.get("/", (req, res) => {
  console.info(`Query:`, req.query.value)
  res.status(200).json({
    message: "All about request object"
  })
})

app.get("/:name", (req, res) => {
  console.info(`Params: `, req.params.name)
  res.status(200).json({
    message: "All about request object"
  })
})

app.post("/", (req, res) => {
  console.info(req.body)
  res.status(200).json({
    message: "All about request object"
  })
})

app.listen("3000", () => {
  console.log("Server is started")
})