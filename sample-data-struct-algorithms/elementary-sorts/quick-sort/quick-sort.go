package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Array struct {
	arr []int
}

func main() {
	var num Array
	num.arr = append(num.arr, 4, 3, 2, 1, 5, 6, 8, 7, 9)
	fmt.Println(num.arr)
	rand.Seed(time.Now().UnixNano())
	//shuffle to prevent quadratic O(n)
	rand.Shuffle(len(num.arr), func(i, j int) { num.arr[i], num.arr[j] = num.arr[j], num.arr[i] })
	num.sort(0, len(num.arr)-1)

	fmt.Println(num.arr)
}

func (a *Array) sort(lo, hi int) {
	if hi <= lo {
		return
	}
	j := a.partition(lo, hi)
	a.sort(lo, j-1)
	a.sort(j+1, hi)
}

func (a *Array) partition(lo, hi int) int {

	i := lo + 1
	j := hi
	V := a.arr[lo]

	for {
		for a.arr[i] < V {
			i++
			if i == hi {
				break
			}
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
