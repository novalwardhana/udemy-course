const db = require("../config/db.config")

module.exports.addPost = (data, callback) => {
  db.query("insert into posts (description, imagePath, dateTimeCreated, addedByUserId) values(?, ?, ?, ?)", 
    [data.description, data.imagePath, data.dateTimeCreated, data.addedByUserId],
    (error, results, fields) => {
      if (error) {
        return callback(new Error(error.message))
      }
      return callback(null, results)
    })
}

module.exports.getAllPosts = (data, callback) => {
  db.query(`
    select 
      p.id as postId, p.description, p.dateTimeCreated, p.likeCount, p.dislikeCount, p.addedByUserId,
      u.firstName, u.lastName
    from posts p
    inner join users as u on p.addedByUserId = u.id
  `,[], (error, results, fields) => {
    if (error) {
      return callback(new Error(error.message))
    }
    return callback(null, results)
  })
}

module.exports.addPostComment = (data, callback) => {
  db.query("insert into comments (postId, comment, dateTimeCreated, addedByUserId) values(?, ?, ?, ?)", 
  [data.postId, data.comment, new Date(), data.addedByUserId],
  (error, results, fields) => {
    if (error) {
      return callback(new Error(error.message))
    }
    return callback(null, results)
  })
}

module.exports.getPostAllComments = (data, callback) => {
  db.query(`
    select 
      c.comment, c.dateTimeCreated, c.addedByUserId, u.firstName, u.lastName
    from comments c
    inner join users as u ON c.addedByUserId = u.id
    where c.postId = ?
  `, [data.postId], (error, results, fields) => {
    if (error) {
      return callback(new Error(error.message))
    }
    return callback(null, results)
  })
}

module.exports.likePost = (data, callback) => {
  db.query(`update posts set likeCount = likeCount + 1 where id = ?`, [data.postId], (error, results, next) => {
    if (error) {
      return callback(new Error(error.message))
    }
    if (results.affectedRows === 1) {
      return callback(null, "Like successful")
    } else {
      return callback(new Error("Invalid post"))
    }
  })
}

module.exports.dislikePost = (data, callback) => {
  db.query(`update posts set dislikeCount = dislikeCount + 1 where id = ?`, [data.postId], (error, results, next) => {
    if (error) {
      return callback(new Error(error.message))
    }
    if (results.affectedRows === 1) {
      return callback(null, "Dislike Successful")
    } else {
      return callback(new Error("Invalid post"))
    }
  })
}

module.exports.deletePost = (data, callback) => {
  db.query(`delete from posts where id = ?`, [data.postId], (error, results, next) => {
    if (error) {
      return callback(new Error(error.message))
    }
    if (results.affectedRows === 1) {
      return callback(null, "Post deleted successfully")
    } else {
      return callback(new Error("Invalid post"))
    }
  })
}