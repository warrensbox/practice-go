package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j >= i; j-- {
			for k := i; k <= j; k++ {
				fmt.Println(arr[k])
			}
			fmt.Println("-----------")
		}
		fmt.Println("+++++")
	}
}
