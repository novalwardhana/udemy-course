package main

import "fmt"

func main() {

	var nilai = 90
	var tugas = 85
	var absensi = 75

	var validasiNilai = nilai > 80
	var validasiTugas = tugas >= 75
	var validasiAbsensi = absensi >= 75

	var result = validasiNilai && validasiTugas && validasiAbsensi
	fmt.Println(result)
	fmt.Println(nilai > 80 || tugas > 85 || absensi > 75)

}
