const express = require("express")
const router = express.Router()
const postsController = require("../controllers/posts.controller")

router.post("/add-post", postsController.addPost)
router.get("/get-all-posts", postsController.getAllPosts)
router.post("/add-post-comment", postsController.addPostComment)
router.get("/get-post-all-comments", postsController.getPostAllComments)
router.put("/like-post", postsController.likePost)
router.put("/dislike-post", postsController.dislikePost)
router.delete("/delete-post", postsController.deletePost)

module.exports = router