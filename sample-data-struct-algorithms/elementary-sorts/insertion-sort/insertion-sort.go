package main

import "fmt"

type Array struct {
	arr []int
}

func main() {

	var num Array
	num.arr = append(num.arr, 1, 3, 2, 4, 6, 8, 5, 7, 0)

	fmt.Println(num.arr)

	for i := 1; i < len(num.arr); i++ {

		for j := i; j > 0 && num.arr[j-1] > num.arr[j]; j-- {
			num.exchange(j, j-1)
			fmt.Println(num.arr)
			fmt.Println("----------")
		}
	}

	fmt.Println(num.arr)
}

func (a *Array) exchange(i, min int) {

	temp := a.arr[i]
	a.arr[i] = a.arr[min]
	a.arr[min] = temp
}
