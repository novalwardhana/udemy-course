package main

import "fmt"

func main() {
	/*
		Membuat slice dari array
		array[low:high]	: Dimulai dari index low sampai high
		array[low:]			: Dimulai dari index low sampai akhir
		array[:high]		: Dimulai dari index 0 sampai high
		array[:]				: Dimulai dari index 0 sampai index terakhir di array

		Slice dan array bersifat pointer => jika kita mengubah elemen turunan array/slice maka element induk/parent juga akan berubah
	*/

	array := [...]string{
		"Januari",
		"februari",
		"Maret",
		"April",
		"Mei",
		"juni",
		"juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(array)
	fmt.Println(len(array))
	fmt.Println(cap(array))

	sliceOne := array[4:9]
	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne)) // 4 sampai 12

	sliceTwo := array[2:3]
	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo))
	fmt.Println(cap(sliceTwo)) // 2 sampai 12

	array[2] = "March"
	sliceOne[3] = "August"
	fmt.Println(array)
	fmt.Println(sliceOne)
	fmt.Println(sliceTwo)

	/* Jika capacity array sudah penuh maka akan membuat variabel reference baru */
	sliceThree := array[10:]
	sliceThree = append(sliceThree, "New Month")
	fmt.Println(sliceThree)
	sliceThree[0] = "Nov"
	fmt.Println(sliceThree)
	fmt.Println(array)

	/* Membuat slice */
	sliceFour := make([]string, 5, 8)
	sliceFour[2] = "Manchester City"
	sliceFour[4] = "Chelsea"
	fmt.Println(sliceFour)
	sliceFour = append(sliceFour, "Arsenal")
	fmt.Println(sliceFour)

	/* Copy SLice */
	sliceFive := make([]string, len(sliceFour), cap(sliceFour))
	copy(sliceFive, sliceFour)
	fmt.Println(sliceFive)

	sliceSix := make([]string, 3, 3)
	copy(sliceSix, sliceFour)
	fmt.Println(sliceSix)
}
