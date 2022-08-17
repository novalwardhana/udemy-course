package main

import (
	"fmt"
)

func TipeDataNumber() {
	/*
		int8		: -128 sampai 127
		int16		: -32768 sampai 32767
		int32		: -2147483648 sampai 2147483647
		int64		: -9223372036854775808 sampai 9223372036854775807

		uint8		: 0 sampai 127
		uint16	: 0 sampai 65535
		uint32	: 0 sampai 4294967295
		uint64	: 0 sampai 18446744073709551615

		float32	: -1.18 x 10^-38 sampai 3.4 x 10^38
		float64	: -2.23 x 10^-308 sampai 1.80 x 10^308

		complex64: complex numbers with float32 real and imaginary parts
		complex128: complex numbers with float64 real and imaginary parts

		byte		: alias dari uint8
		rune		: alias dari int32
		int 		: minimal int32
		uint		: minimal uint32

	*/
	fmt.Println("Int: ", 1)
	fmt.Println("float64: ", 3.60)
}
