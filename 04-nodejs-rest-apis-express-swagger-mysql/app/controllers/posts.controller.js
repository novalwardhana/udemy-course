const postsService = require("../services/posts.service")

module.exports.addPost = (req, res, next) => {
  const data = {
    description: req.body.description,
    imagePath: req.body.imagePath,
    dateTimeCreated: new Date().toISOString().slice(0, 19).replace("T", " "),
    addedByUserId: req.body.addedByUserId,
  }
  postsService.addPost(data, (error, results) => {
    if (error) {
      console.info(error)
      return res.status(400).send({ success: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}

module.exports.getAllPosts = (req, res, next) => {
  const data = {}
  postsService.getAllPosts(data, (error, results) => {
    if (error) {
      return res.status(400).send({ status: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}

module.exports.addPostComment = (req, res, next) => {
  const data = {
    postId: req.body.postId,
    comment: req.body.comment,
    addedByUserId: req.body.addedByUserId
  }
  postsService.addPostComment(data, (error, results) => {
    if (error) {
      return res.status(200).send({ status: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}

module.exports.getPostAllComments = (req, res, next) => {
  const data = {
    postId: req.query.postId
  }
  postsService.getPostAllComments(data, (error, results) => {
    if (error) {
      return res.status(400).send({ status: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}

module.exports.likePost = (req, res, next) => {
  const data = {
    postId: req.body.postId
  }
  postsService.likePost(data, (error, results) => {
    if (error) {
      return res.status(400).send({ status: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}

module.exports.dislikePost = (req, res, next) => {
  const data = {
    postId: req.body.postId
  }
  postsService.dislikePost(data, (error, results) => {
    if (error) {
      return res.status(400).send({ status: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}

module.exports.deletePost = (req, res) => {
  const data = {
    postId: req.query.postId
  }
  postsService.deletePost(data, (error, results) => {
    if (error) {
      return res.status(400).send({ status: 0, data: error.message })
    }
    return res.status(200).send({
      status: 1,
      data: results
    })
  })
}