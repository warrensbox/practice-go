package main

import "fmt"

type Array struct {
	arr []int
}

func main() {

	var num Array
	num.arr = append(num.arr, 1, 3, 2, 4, 6, 8, 5, 7, 2)

	fmt.Println(num.arr)

	for i := 0; i < len(num.arr); i++ {
		min := i
		for j := i + 1; j < len(num.arr); j++ {
			if num.arr[min] > num.arr[j] {
				num.exchange(i, j)
			}
		}
	}

	fmt.Println(num.arr)
}

func (a *Array) exchange(i, min int) {

	temp := a.arr[i]
	a.arr[i] = a.arr[min]
	a.arr[min] = temp

}
