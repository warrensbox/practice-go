package main

import (
	"fmt"
)

type Array struct {
	arr []int
}

func main() {
	var num Array
	num.arr = append(num.arr, 7, 3, 2, 8, 5, 6, 1, 4, 9)
	//fmt.Println(num.arr)
	//rand.Seed(time.Now().UnixNano())
	//shuffle to prevent quadratic O(n^2)
	//rand.Shuffle(len(num.arr), func(i, j int) { num.arr[i], num.arr[j] = num.arr[j], num.arr[i] })
	//	num.arr[0], num.arr[3] = num.arr[3], num.arr[0]
	num.sort(0, len(num.arr)-1)

	fmt.Println(num.arr)
}

func (a *Array) sort(lo, hi int) {
	fmt.Println("HI", hi)
	fmt.Println("LO", lo)
	if hi <= lo {
		return
	}
	j := a.partition(lo, hi)
	fmt.Println("JJ", j)
	a.sort(lo, j-1)
	fmt.Println("j", j)
	fmt.Println("hi", hi)
	a.sort(j+1, hi)
}

func (a *Array) partition(lo, hi int) int {

	i := lo + 1
	j := hi
	V := a.arr[lo]

	fmt.Println(a.arr)
	for {
		for a.arr[i] < V {
			i++
			if i == hi {
				break
			}
		}
		if a.arr[j] == 2 {
			fmt.Println("here")
			fmt.Println("j", j)
			fmt.Println("i", i)
		}
		for a.arr[j] > V {
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a.exchange(i, j)
	}
	a.exchange(lo, j)
	return j
}

func (a *Array) exchange(i, min int) {
	a.arr[i], a.arr[min] = a.arr[min], a.arr[i]
}
