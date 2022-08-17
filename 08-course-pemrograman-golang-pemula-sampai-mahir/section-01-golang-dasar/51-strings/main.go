package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Novalita Kusuma Wardhana", "Kusuma"))
	fmt.Println(strings.Contains("Novalita Kusuma", "ta Ku"))
	fmt.Println(strings.Contains("Novalita kusuma", "Wardhana"))

	fmt.Println(strings.Split("Novalita Kusuma Wardhana", " "))
	fmt.Println(strings.Split("a,b,c,d,e,f,g,h,i", ","))

	fmt.Println(strings.ToUpper("noVAliTa"))
	fmt.Println(strings.ToUpper("arSeNAl"))
	fmt.Println(strings.ToLower("KUsuMa"))
	fmt.Println(strings.ToLower("CHelSEa"))
	fmt.Println(strings.ToTitle("warDHana"))
	fmt.Println(strings.ToTitle("liVERpooL"))

	fmt.Println(strings.Trim("     Novalita Kusuma Wardhana   ", " "))

	fmt.Println(strings.ReplaceAll("Noval Kusuma Noval Wardhana", "Noval", "Val"))
	fmt.Println(strings.Replace("Noval Kusuma Noval Wardhana", "Noval", "Val", 1))

	fmt.Println(strings.Join([]string{"Arsenal", "Chelsea", "Liverpool", "Barcelona"}, "-"))
	fmt.Println(strings.Join([]string{"Kota Yogyakarta", "Bantul", "Sleman", "Kulon Progo", "Gunung Kidul"}, " "))
}
