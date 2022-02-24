/*
  URL
  a. Standard library untuk memanipulasi URL
  b. Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/url.html
*/

import { url } from "inspector"
import {URL} from "url"

/* basic */
const url1 = new URL("https://novalwardhana.com/about/me?age=20")
console.info(`url1 href: `, url1.href)
console.info(`url1 toString: `, url1.toString())
console.info(`url1 protocol: `, url1.protocol)
console.info(`url1 host: `, url1.host)
console.info(`url1 hostname: `, url1.hostname)
console.info(`url1 search params: `, url1.searchParams)

/* Manipulasi URL */
url1.protocol = "wss"
url1.host = "noval.co.id"
url1.searchParams.append("address", "bantul")
console.info(`url1 href: `, url1.href)
console.info(`url1 toString: `, url1.toString())
console.info(`url1 protocol: `, url1.protocol)
console.info(`url1 host: `, url1.host)
console.info(`url1 hostname: `, url1.hostname)
console.info(`url1 search params: `, url1.searchParams)

const url2 = new URL("https://hipwee.com/story/user?name=noval&from=bantul")
console.info(`url2: ${url2.toString()}`)
console.info(`url2 href: ${url2.href}`)
console.info(`url2 query params: ${url2.searchParams}`)
console.info(`url2 host: ${url2.host}`)
console.info(`url2 pathname: ${url2.pathname}`)
console.info(`url2 protocol: ${url2.protocol}`)
url2.host = "hipwee.co.id"
url2.searchParams.delete("name")
url2.searchParams.delete("from")
url2.searchParams.append("age", 25)
url2.searchParams.append("city", "sleman")
console.info(`url2 now: ${url2.href}`)