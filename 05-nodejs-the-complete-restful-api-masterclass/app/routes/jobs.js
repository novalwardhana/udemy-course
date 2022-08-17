const express = require("express")
const router = express.Router()
const { 
  getJobs, newJob, getJobsInRadius, updateJob, deleteJob, getJob, jobStats
} = require("../controller/jobsController")

router.route("/jobs").get(getJobs)
router.route("/job/new").post(newJob)
router.route("/jobs/:zipcode/:distance").get(getJobsInRadius)
router.route("/job/:id").put(updateJob)
router.route("/job/:id").delete(deleteJob)
router.route("/job/:id/:slug").get(getJob)
router.route("/stats/:topic").get(jobStats)

module.exports = router