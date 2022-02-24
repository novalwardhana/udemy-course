import fs from "fs"
import zlib from "zlib"

const file = fs.readFileSync("compress-file.txt")
const fileDecompress = zlib.unzipSync(file)
fs.writeFileSync("decompress-file.txt", fileDecompress)

const file2 = fs.readFileSync("compress-file2.txt")
const fileDecompress2 = zlib.unzipSync(file2)
fs.writeFileSync("decompress-file2.txt", fileDecompress2)