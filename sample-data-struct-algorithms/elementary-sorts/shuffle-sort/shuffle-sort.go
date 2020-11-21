package main

import (
	"fmt"
	"math/rand"
)

type Array struct {
	arr []int
}

func main() {

	var num Array
	num.arr = append(num.arr, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(num.arr)

	// rand.Intn returns a random int n, 0 <= n < x
	for i := 0; i < len(num.arr); i++ {
		r := rand.Intn(i + 1)
		num.exchange(i, r)
	}

	fmt.Println(num.arr)
}

func (a *Array) exchange(i, min int) {
	a.arr[i], a.arr[min] = a.arr[min], a.arr[i]
}
