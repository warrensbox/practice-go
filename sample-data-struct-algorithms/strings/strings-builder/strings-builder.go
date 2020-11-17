package main

import (
	"fmt"
	"strings"
)

func main() {

	var result strings.Builder

	str := "hello word"

	c := strings.Fields(str)

	for _, char := range c {
		fmt.Println(char)
		result.WriteString(char + "a") //will append
	}

	temp := result.String()

	fmt.Println(temp)
}
