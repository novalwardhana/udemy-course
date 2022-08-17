package main

import (
	"46-package-import/helper"
)

func main() {
	/*
		Package
		1. Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di golang
		2. Dengan menggunakan package kita bisa merapikan kode program yang kita buat
		3. Package sendiri hanyalah direktori folder di sistem operasi kita

		Import
		1. Secara standar, fila golang hanya bisa mengakses file golang lainnya yang berada dalam package yang sama
		2. Jika kita ingin mengakses file golang yang berada diluar package, maka kita bisa menggunakan import
	*/
	helper.SayHello("Noval")
	helper.SayHai("Wardhana")
}
