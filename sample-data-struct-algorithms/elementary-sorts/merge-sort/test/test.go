package main

import (
	"fmt"
)

func main() {

	var num Array
	num.arr = append(num.arr, 4, 3, 2, 1, 5, 6, 8, 7)
	num.aux = make([]int, len(num.arr))

	num.sort(0, len(num.arr)-1)

	fmt.Println(num.arr)
}

type Array struct {
	arr []int
	aux []int
}

func (a *Array) sort(lo, hi int) {
	fmt.Println("=========")
	fmt.Println("lo", lo)
	fmt.Println("hi", hi)

	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	a.sort(lo, mid)
	a.sort(mid+1, hi)
	merge(lo, mid, hi)
}

func merge(lo, mid, hi int) {
	fmt.Println()
	fmt.Println("merge")
	fmt.Println("LO", lo)
	fmt.Println("MID", mid)
	fmt.Println("HI", hi)

}
