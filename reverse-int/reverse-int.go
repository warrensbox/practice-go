package main

import "fmt"

func main(){

	ans := reverse(123456)

	fmt.Println(ans)
}

func reverse(x int) int {

	newint := 0
	val := 0

	for x > 0 {
		remains := x%10
		val = remains+newint
		newint = val*10
		x = x/10
	}

	return val    
}
