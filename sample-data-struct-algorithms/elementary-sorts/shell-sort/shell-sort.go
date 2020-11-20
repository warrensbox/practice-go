package main

import (
	"fmt"
)

type Array struct {
	arr []int
}

func main() {

	index := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	var num Array
	num.arr = append(num.arr, 1, 3, 2, 4, 6, 8, 5, 7, 0)
	num.arr = append(num.arr, 12, 15, 14, 16, 18, 17, 19, 21, 20, 22, 23)
	num.arr = append(num.arr, 24, 15, 17, 16, 18, 19, 31, 30, 32, 33, 36, 34, 35)
	fmt.Println(index)
	fmt.Println(num.arr)

	n := len(num.arr)
	h := 1

	for h < n/3 {
		h = 3*h + 1 //4,13,40
	}

	for h >= 1 { //h sort array
		for i := h; i < len(num.arr); i++ { //insertion sort

			for j := i; j >= h && num.arr[j-h] > num.arr[j]; j = j - h {
				num.exchange(j, j-h)
				fmt.Println(num.arr)
				fmt.Println("----------")
			}
		}
		h = (h - 1) / 3
	}

	fmt.Println(num.arr)
}

func (a *Array) exchange(i, min int) {

	temp := a.arr[i]
	a.arr[i] = a.arr[min]
	a.arr[min] = temp
}
