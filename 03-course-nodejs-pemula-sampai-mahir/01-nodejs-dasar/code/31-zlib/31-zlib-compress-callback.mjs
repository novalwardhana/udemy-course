import fs from "fs"
import zlib from "zlib"

fs.readFile("sample.txt", (error, data) => {
  if (error !== null) {
    console.info(error.message)
    return
  }
  zlib.gzip(data, (error, compressData) => {
    if (error !== null) {
      console.info(error.message)
      return
    }
    fs.writeFile("compress-file3.txt", compressData, (error) => {
      if (error !== null) {
        console.info(error.message)
        return
      }
      console.info("success compress file")
    })
  })
})