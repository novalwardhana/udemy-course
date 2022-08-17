package main

import "fmt"

func TypeDeclarations() {
	type Identity string
	type Status bool

	var KTP Identity = "34021220123321456654"
	var Paspor Identity = "78934278"
	var marriageStatus Status = false
	var EmployeStatus Status = true
	fmt.Println(KTP)
	fmt.Println(Paspor)
	fmt.Println(marriageStatus)
	fmt.Println(EmployeStatus)
}
