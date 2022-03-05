const usersService = require("../services/users.service")

module.exports.login = (req, res, next) => {
  const data = {
    emailId: req.body.emailId,
    password: req.body.password
  }
  usersService.login(data, (error, results) => {
    console.info("error: ", error)
    console.info("results: ", results)
    if (error) {
      return res.status(404).send({ success: 0, data: error.message })
    }
    return res.status(200).send({
      success: 1,
      data: results
    })
  })
}

module.exports.register = (req, res, next) => {
  const data = {
    firstName: req.body.firstName,
    lastName: req.body.lastName,
    emailId: req.body.emailId,
    password: req.body.password
  }
  console.info("ini data", data)
  usersService.register(data, (error, results) => {
    if (error) {
      return res.status(400).send({ success: 0, data: "Bad request" })
    }
    return res.status(201).send({
      success: 1,
      data: "Registration successfully"
    })
  })
}