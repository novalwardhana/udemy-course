module.exports = (err, req, res, next) => {
  err.statusCode = err.statusCode || 500

  if (process.env.NODE_ENV === "development") {
    return res.status(err.statusCode).json({
      status: false,
      error: err,
      errMessage: err.message,
      errStack: err.stack
    })
  }

  err.message = err.message || "Internal server error"
  if (process.env.NODE_ENV === "production") {
    let error = {...err}
    error.message = err.message
    return res.status(err.statusCode).json({
      success: false,
      message: error.message || "Internal server error"
    })
  }

  return res.status(err.statusCode).json({
    status: false,
    message: err.message
  })
}