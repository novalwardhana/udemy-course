import net from "net"

const server = net.createServer((client) => {
  client.addListener("data", (data) => {
    console.info(`Receiving from client: ${data.toString()}`)
    client.write(`${data.toString()}\r\n`)
  })
})
server.listen("3030", "localhost")