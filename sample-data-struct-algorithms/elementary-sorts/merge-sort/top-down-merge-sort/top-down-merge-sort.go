package main

import "fmt"

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

	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	a.sort(lo, mid)
	a.sort(mid+1, hi)
	a.merge(lo, mid, hi)
}

func (a *Array) merge(lo, mid, hi int) {

	i := lo
	j := mid + 1

	fmt.Println("aux", a.aux)
	fmt.Println("arr", a.arr)
	for k := lo; k <= hi; k++ { //== because it is not the length of
		a.aux[k] = a.arr[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			a.arr[k] = a.aux[j]
			j++
		} else if j > hi {
			a.arr[k] = a.aux[i]
			i++
		} else if a.aux[j] < a.aux[i] {
			a.arr[k] = a.aux[j]
			j++
		} else {
			a.arr[k] = a.aux[i]
			i++
		}

	}

	fmt.Println("=====")

}
