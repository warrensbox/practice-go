package main

import (
	"fmt"
	"strings"
)

func main() {

	//input
	S := "2-5g-3-J"
	K := 2

	output := licenseKeyFormatting(S, K)
	fmt.Println(output)
}

func licenseKeyFormatting(S string, K int) string {

	S = strings.Replace(S, "-", "", -1)
	S = strings.ToUpper(S)

	modulus := len(S) % K
	fmt.Println("modulus", modulus)

	if modulus == 0 {
		modulus += K
	}

	fmt.Println("modulus", modulus)
	fmt.Println("length", len(S))

	for modulus < len(S) {
		S = S[:modulus] + "-" + S[modulus:]
		fmt.Println("S", S)
		modulus += K + 1
		fmt.Println("modulus", modulus)
		fmt.Println("---")
	}

	return S
}
