package main

import "fmt"

func main() {
	/*
		Defer
		1. Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
		2. Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	*/
	runApplication(10)
	runApplication(3)
	runApplication(5)

	/*
		Panic
		1. Function yang bisa kita gunakan untuk menghentikan program
		2. Panic biasanya dipanggil ketika terjadi error pada saat program kita berjalan
		3. Saat panic program akan berhenti tapi defer akan tetap dieksekusi
	*/
	runPanic(false)
	runPanic(false)

	/*
		Recover
		1. Recover adalah function yang bisa kita gunakan untuk menangkap data panic
		2. Dengan recover proses panic akan terhenti dan program akan tetap berjalan
	*/
	runRecover(true)
	fmt.Println("Test recover in main function")
	runRecover(false)
	runRecover(true)
}

func loggingOne() {
	fmt.Println("Selesai memanggil function logging 1")
}

func loggingTwo() {
	fmt.Println("Selesai memanggil function logging 2")
}

func runApplication(value int) {
	defer loggingTwo()
	defer loggingOne()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println(result)
}

func runPanic(check bool) {
	defer loggingTwo()
	defer loggingOne()
	if check {
		panic("Test panic")
	}
	fmt.Println("Run application")
}

func runRecover(check bool) {
	defer loggingTwo()
	defer loggingOne()
	defer func() {
		message := recover()
		if message != nil {
			fmt.Println("Test panic message: ", message)
		}
	}()
	if check {
		panic("Test panic recover")
	}
	fmt.Println("Run recover")
}
