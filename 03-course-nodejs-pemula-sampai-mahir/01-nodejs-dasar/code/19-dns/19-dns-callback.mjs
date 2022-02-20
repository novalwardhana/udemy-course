/*
  DNS
  DNS merupakan standard library yang bisa digunakan untuk bekerja dengan DNS (Domain Name Server)
  Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/dns.html
*/

import dns from "dns"

dns.lookup("www.programmerzamannow.com", (error, address, family) => {
  if (error !== null) {
    console.info(`Cannot get DNS1: `, error.message)
    return
  }
  console.info(`DNS1 address: ${address}, family: ${family}`)
})

dns.lookup("www.goooooogle.co.id", (error, address, family) => {
  if (error != null) {
    console.info(`Cannot get DNS2: `, error.message)
    return
  }
  console.info(`DNS2 address: ${address}, family: ${family}`)
})

dns.lookup("www.novalwardhana.com", (error, address, family) => {
  if (error != null) {
    console.info(`Cannot get DNS3: ${error.message}`)
    return
  }
  console.info(`DNS3 address: ${address}, family: ${family}`)
})