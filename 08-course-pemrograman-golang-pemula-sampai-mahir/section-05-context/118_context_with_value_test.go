package belajar_golang_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	/* Set value */
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")

	contextH := context.WithValue(contextD, "h", "H")
	contextI := context.WithValue(contextD, "i", "I")

	fmt.Println("contextA:", contextA)
	fmt.Println("contextB:", contextB)
	fmt.Println("contextC:", contextC)
	fmt.Println("contextD:", contextD)
	fmt.Println("contextE:", contextE)
	fmt.Println("contextF:", contextF)
	fmt.Println("contextG:", contextG)
	fmt.Println("contextH:", contextH)
	fmt.Println("contextI:", contextI)
	fmt.Printf("-----\n\n")

	/* Get value: bisa mendapat data dari contextnya sendiri */
	fmt.Println(contextB.Value("b"))
	fmt.Println(contextC.Value("c"))
	fmt.Println(contextD.Value("d"))
	fmt.Println(contextE.Value("e"))
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextG.Value("g"))
	fmt.Println(contextH.Value("h"))
	fmt.Println(contextI.Value("i"))
	fmt.Printf("-----\n\n")

	/* Get value: parent tidak bisa mendapat value dari child */
	fmt.Println(contextA.Value("b"))
	fmt.Println(contextB.Value("d"))
	fmt.Println(contextB.Value("e"))
	fmt.Printf("-----\n\n")

	/* Get value: child bisa mendapat data dari parent */
	fmt.Println(contextD.Value("b"))
	fmt.Println(contextE.Value("b"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextG.Value("c"))
	fmt.Println(contextH.Value("b"))
	fmt.Printf("-----\n\n")

	/* Get value: child tidak bisa mendapat data dari parent yang berbeda */
	fmt.Println(contextD.Value("c"))
	fmt.Println(contextE.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextG.Value("b"))
	fmt.Println(contextI.Value("c"))
	fmt.Printf("-----\n\n")
}

type ContextWithValueCardbased struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ContextWithValueEwallet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestContextWithValueStruct(t *testing.T) {
	ElectronicTrx := context.Background()

	/* Group 1 */
	ElectronicTrxEwallet := context.WithValue(ElectronicTrx, "ewallet", "E-Wallet")
	ElectronicTrxEwalletOVO := context.WithValue(ElectronicTrxEwallet, "ovo", &ContextWithValueEwallet{
		ID:   1,
		Name: "OVO Wallet",
	})
	ElectronicTrxWalletGopay := context.WithValue(ElectronicTrxEwallet, "gopay", ContextWithValueEwallet{
		ID:   2,
		Name: "Gopay",
	})
	ElectronicTrxWalletDana := context.WithValue(ElectronicTrxEwallet, "dana", &ContextWithValueEwallet{
		ID:   3,
		Name: "Dana",
	})
	fmt.Println("Get value: self context")
	fmt.Println(ElectronicTrxEwallet.Value("ewallet"))
	fmt.Println(ElectronicTrxEwalletOVO.Value("ovo"))
	fmt.Println(ElectronicTrxWalletGopay.Value("gopay"))
	fmt.Println(ElectronicTrxWalletDana.Value("dana"))
	fmt.Printf("-----\n\n")

	/* Group 2 */
	ElectronicTrxChipbased := context.WithValue(ElectronicTrx, "chipbased", "Chipbased")
	ElectronicTrxChipbasedEmoney := context.WithValue(ElectronicTrxChipbased, "emoney", &ContextWithValueCardbased{
		ID:   1,
		Name: "Mandiri E-Money",
	})
	ElectronicTrxChipbasedBrizzi := context.WithValue(ElectronicTrxChipbased, "brizzi", ContextWithValueCardbased{
		ID:   2,
		Name: "BCA Brizzi",
	})
	fmt.Println("Get value: from parent")
	fmt.Println(ElectronicTrxChipbasedEmoney.Value("chipbased"))
	fmt.Println(ElectronicTrxChipbasedBrizzi.Value("chipbased"))
	fmt.Printf("-----\n\n")

	fmt.Println("Get value: from child")
	fmt.Println(ElectronicTrx.Value("ewallet"))
	fmt.Println(ElectronicTrx.Value("chipbased"))
	fmt.Println(ElectronicTrx.Value("brizzi"))
	fmt.Printf("-----\n\n")
}
