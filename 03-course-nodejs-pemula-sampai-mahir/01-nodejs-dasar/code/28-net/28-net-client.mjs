import net from "net"

const client = net.createConnection({
  port: 3000,
  host: "localhost"
})
client.addListener("data", (data) => {
  console.info(`Receiving data from server: ${data.toString()}`)
})

setInterval(() => {
  client.write(`Noval ${process.argv[2]}\r\n  `)
}, 2000)