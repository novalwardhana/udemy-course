package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Strconv
		Untuk melakukan konversi dari string ke int dan sebaliknya
	*/

	/* Parse bool */
	bool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("err parse bool: ", err)
	} else {
		fmt.Println("success parse bool: ", bool)
	}
	bool, err = strconv.ParseBool("iniTrue")
	if err != nil {
		fmt.Println("err parse bool: ", err)
	} else {
		fmt.Println("success parse bool: ", bool)
	}
	bool, err = strconv.ParseBool("false")
	if err != nil {
		fmt.Println("err parse bool: ", err)
	} else {
		fmt.Println("success parse bool: ", bool)
	}

	/* Parse Int */
	integer, err := strconv.ParseInt("10000", 10, 32)
	if err != nil {
		fmt.Println("Err parse int: ", err)
	} else {
		fmt.Println("success parse int: ", integer)
	}
	integer, err = strconv.ParseInt("10000.5", 10, 64)
	if err != nil {
		fmt.Println("Err parse int: ", err)
	} else {
		fmt.Println("Success parse int: ", integer)
	}
	integer, err = strconv.ParseInt("125", 10, 64)
	if err != nil {
		fmt.Println("Err parse int: ", err)
	} else {
		fmt.Println("success parse int: ", integer)
	}

	/* Parse float */
	float, err := strconv.ParseFloat("12.56", 64)
	if err != nil {
		fmt.Println("Err parse float: ", err)
	} else {
		fmt.Println("Success parse float: ", float)
	}
	float, err = strconv.ParseFloat("85", 32)
	if err != nil {
		fmt.Println("Err parse float: ", err)
	} else {
		fmt.Println("Success parse float: ", float)
	}
	float, err = strconv.ParseFloat("a12b", 64)
	if err != nil {
		fmt.Println("Err parse float: ", err)
	} else {
		fmt.Println("Success parse float: ", float)
	}

	/* Format Int */
	formatBool := strconv.FormatBool(true)
	fmt.Println("Success format bool: ", formatBool)

	/* Format Int */
	formatInt := strconv.FormatInt(20, 2)
	fmt.Println("Success format int: ", formatInt)
	formatInt = strconv.FormatInt(20, 6)
	fmt.Println("Success format int: ", formatInt)
	formatInt = strconv.FormatInt(20, 16)
	fmt.Println("Success format int: ", formatInt)
	formatInt = strconv.FormatInt(20, 10)
	fmt.Println("Success format int: ", formatInt)

	/* ATOI */
	atoi, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error atoi: ", err)
	} else {
		fmt.Println("Success atoi: ", atoi)
	}
	atoi, err = strconv.Atoi("1a2b3c")
	if err != nil {
		fmt.Println("Error atoi: ", err)
	} else {
		fmt.Println("Success atoi: ", atoi)
	}

	/* ITOA */
	itoa := strconv.Itoa(200)
	fmt.Println("Success itoa: ", itoa)
	itoa = strconv.Itoa(3000)
	fmt.Println("Success itoa: ", itoa)
}
