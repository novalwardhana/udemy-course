const mongoose = require("mongoose")
const validator = require("validator")
const slugify = require("slugify")
const geoCoder = require("../utils/geocoder")

const jobsSchema = new mongoose.Schema({
  title: {
    type: String,
    required: [true, "Please enter job title"],
    trim: true,
    maxlength: [100, "Job title cannot exceed 100 characters..."]
  },
  slug : String,
  description: {
    type: String,
    required: [true, "Please enter job description"],
    maxlength: [1000, "Job description cannot exceed 1000 characters..."]
  },
  email: {
    type: String,
    validate: [validator.isEmail, "Please add valid email..."]
  },
  address: {
    type: String,
    required: [true, "Please add an address..."]
  },
  location: {
    type: {
      type: String,
      enum: [
        "Point"
      ]
    },
    coordinates: {
      type: [Number],
      index: "2dsphere"
    },
    formattedAddress: String,
    city: String,
    state: String,
    zipcode: String,
    country: String
  },
  company: {
    type: String,
    required: [true, "Please add company name..."]
  },
  industry: {
    type: [String],
    required: true,
    enum: {
      values: [
        'Bussiness',
        'Information Technology',
        'Banking',
        'Education/Training',
        'Others'
      ],
      message: "Please select correct options for industry"
    }
  },
  jobType: {
    type: String,
    required: [true, "Please enter job type"],
    enum: {
      values: [
        "Permanent",
        "Temporary",
        "Inthernship"
      ],
      message: "Please select correct options for job type..."
    }
  },
  minEducation: {
    type: String,
    required: [true, "Please enter min education"],
    enum: {
      values: [
        "Bachelors",
        "Masters",
        "Phd"
      ],
      message: "Please select correct option for education..."
    }
  },
  positions: {
    type: Number,
    default: 1
  },
  experience: {
    type: String,
    required: [true, "Please enter experience"], 
    enum: {
      values: [
        "No Experience",
        "1 Year - 2 Years",
        "2 Year - 5 Years",
        "5 Years+"
      ],
      message: "PLease select orrect options for experience"
    }
  },
  salary: {
    type: Number,
    required: [true, "Please insert expected salary for this jobs..."]
  },
  postingDate: {
    type: Date,
    default: Date.now()
  },
  lastDate: {
    type: Date,
    default: new Date().setDate(new Date().getDate() + 7)
  },
  applicantsApplied: {
    type: [Object],
    select: false
  }
})

jobsSchema.pre("save", function(next) {
  this.slug = slugify(this.title, {lower: true})

  next()
})

jobsSchema.pre("save", async function(next) {
  const loc = await geoCoder.geocode(this.address)

  this.location = {
    type: "Point",
    coordinates: [loc[0].longitude, loc[0].latitude],
    formattedAddress: loc[0].formattedAddress,
    city: loc[0].address,
    state: loc[0].stateCode,
    zipcode: loc[0].zipcode,
    country: loc[0].countryCode
  }
})

module.exports = mongoose.model("Job", jobsSchema)