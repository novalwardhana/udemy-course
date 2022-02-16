/*
  NodeJS Standard Library
  a. Saat kita belajar javascript di web browser, terdapat fitur-fitur bernama WEB-API
    referensi: https://developer.mozilla.org/en-US/docs/Web/API
  b. Kebanyakan fitur WEB-API hanya berjalan di browser, sehingga tidak bisa jalan di NodeJS
  c. NodeJS sendiri hanya menggunakan bahasa pemrograman Javascript, namun tidak mengadopsi fitur WEB-API
    karena itu hanya berjalan di web browser
  d. NodeJS sendiri memiliki standard library yang bisa kita gunakan untuk mempermudah pembuatan aplikasi
    Referensi: https://nodejs.org/dist/latest-v16.x/docs/api/ 
*/

const request = new XMLHttpRequest() // Akan error karena di nodejs tidak ada library XMLHttpRequest