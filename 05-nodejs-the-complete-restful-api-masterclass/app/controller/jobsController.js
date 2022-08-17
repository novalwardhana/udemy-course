const Job = require("../models/jobs")
const geoCoder = require("../utils/geocoder")
const ErrorHandler = require("../utils/errorHandler")
const catchAsyncErrors = require("../middlewares/catchAsyncErrors")

/* Get all jobs */
exports.getJobs = catchAsyncErrors( async (req, res, next) => {
  const job = await  Job.find()
  res.status(200).json({
    status: true,
    results: job.length,
    data: job
  })
})

/* Create new JObs */
exports.newJob = catchAsyncErrors( async (req, res, next) => {
  const job = await Job.create(req.body)
  res.status(200).json({
    success: true,
    message: "Job created",
    data: job
  })
})

/* Search jobs with radius */
exports.getJobsInRadius = catchAsyncErrors (async (req, res, next) => {
  const {zipcode, distance} = req.params

  const loc = await geoCoder.geocode(zipcode)
  const latitude = loc[0].latitude
  const longitude = loc[0].longitude

  const radius = distance / 3963

  const jobs = await Job.find({
    location: {$geoWithin: {$centerSphere: [[longitude, latitude], radius]}}
  })

  res.status(200).json({
    success: true, 
    results: jobs.length,
    data: jobs
  })
})

/* Update job */
exports.updateJob = catchAsyncErrors( async  (req, res, next) => {
  let job = await Job.findById(req.params.id)

  if (!job) {
    return next(new ErrorHandler("job not found", 404))
  }

  job = await Job.findByIdAndUpdate(req.params.id, req.body, {
    new: true,
    runValidators: true
  })

  res.status(200).json({
    status: true, 
    message: "Job is updated",
    data: job
  })
})

/* Delete job */
exports.deleteJob = catchAsyncErrors( async (req, res, next) => {
  let job = await jobs.findById(req.params.id)

  if (!job) {
    return res.status(404).json({
      status: false,
      message: "Job not found"
    })
  }

  job = await jobs.findByIdAndDelete(req.params.id)

  res.status(200).json({
    status: true,
    message: "Success delete job"
  })
})

/* get single job by id */
exports.getJob = catchAsyncErrors( async (req, res, next) => {
  let job = await Job.find({$and: [{_id: req.params.id}, {slug: req.params.slug}]})

  if (!job || job.length === 0) {
    return res.status(404).json({
      status: false,
      message: "Job not found"
    })
  }

  res.status(200).json({
    status: true, 
    data: job
  })
})

/* 
  get jon stats 
  Must create index in mongoDB:  db.jobs.createIndex({title : "text"})
*/
exports.jobStats = catchAsyncErrors( async (req, res, next) => {
  const stats = await Job.aggregate([
    {
      $match: {$text : {$search : "\""+ req.params.topic + "\"" }}
    },
    {
      $group: {
        _id: {$toUpper: '$experience'},
        totalJobs: {$sum: 1},
        avgPosition : {$avg: '$positions'},
        avgSalary : {$avg : '$salary'},
        minSalary: {$min : '$salary'},
        maxSalary: {$max: '$salary'}
      }
    }
  ])

  if (stats.length === 0) {
    return res.status(200).json({
      success: false,
      message: "Data not found"
    })
  }

  res.status(200).json({
    success: true,
    data: stats
  })
})