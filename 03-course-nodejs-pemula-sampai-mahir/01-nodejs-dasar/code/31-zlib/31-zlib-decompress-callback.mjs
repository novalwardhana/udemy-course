import fs from "fs"
import zlib from "zlib"

fs.readFile("compress-file3.txt", (error, data) => {
  if (error !== null) {
    console.info(error.message)
    return
  }
  zlib.unzip(data, (error, dataDecompress) => {
    if (error !== null) {
      console.info(error.message)
      return
    }
    fs.writeFile("decompress-file3.txt", dataDecompress, (error) => {
      if (error !== null) {
        console.info(error.message)
        return
      }
      console.info("Success decompress file")
    })
  })
})
console.info("Start")