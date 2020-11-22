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
	lt := lo
	i := lo + 1
	gt := hi
	v := a.arr[lo]

	for i <= gt {
		cmp := a.compareTo(i, v)
		if cmp < 0 {
			a.exchange(lt, i)
			lt++
			i++
		} else if cmp > 0 {
			a.exchange(i, gt)
			gt--
		} else {
			i++
		}
	}

	a.sort(lo, lt-1)
	a.sort(gt+1, hi)
}

func (a *Array) compareTo(i, v int) int {
	if a.arr[i] > v {
		return 1
	} else if a.arr[i] < v {
		return -1
	}

	return 0
}

func (a *Array) exchange(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}
