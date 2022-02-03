package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "chicken"

	idx := strings.Index(str, "ken")

	fmt.Print(str[:idx])

}
