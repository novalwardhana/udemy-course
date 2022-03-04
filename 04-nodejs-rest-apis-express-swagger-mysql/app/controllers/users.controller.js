const usersService = require("../services/users.service")

module.exports.register = (req, res, next) => {
  console.info("req", req.body)
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
      data: results
    })
  })
}