/*
  DNS
  DNS merupakan standard library yang bisa digunakan untuk bekerja dengan DNS (Domain Name Server)
  Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/dns.html
*/

import dns from "dns/promises"

const dns1 = await dns.lookup("www.programmerzamannow.com")
console.info(`DNS1 address: ${dns1.address}, family: ${dns1.family}`)

const dns2 = await dns.lookup("www.google.co.id")
console.info(`DNS2 address: ${dns2.address}, family: ${dns2.family}`)

const dns3 = await dns.lookup("www.novalwardhana.com")
console.info(`DNS3 address: ${dns3.address}, family: ${dns3.family}`)


