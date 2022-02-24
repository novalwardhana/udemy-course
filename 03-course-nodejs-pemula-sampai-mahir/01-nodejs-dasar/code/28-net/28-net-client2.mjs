import net from "net"

const client = net.createConnection({
  port: 3030,
  host: "localhost"
})
client.addListener("data", (data) => {
  console.info(`Receiving from server: ${data.toString()}`)
})

setInterval(() => {
  client.write(`Test client server\r\n`)
}, 2000)