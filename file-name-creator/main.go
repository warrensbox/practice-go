package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	arg := os.Args[1]

	if arg == "" {
		fmt.Println("nothing")
		return
	}

	str := strings.Replace(arg, " ", "-", -1)

	fmt.Println(strings.ToLower(str))
}
