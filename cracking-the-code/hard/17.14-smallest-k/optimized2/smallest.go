package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	k := 7 //the smallest 7 elements
	arr := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	num := Array{}
	num.arr = arr

	num.Sort(0, len(arr)-1, k)
	//fmt.Println(num.arr)
}

type Array struct {
	arr []int
}

func (a *Array) Sort(lo, hi, m int) {

	if hi <= lo {
		return
	}

	j := a.Partition(lo, hi)
	if j == m {
		a.Print(j)
		return
	} else if j < m {
		lo = j + 1
	} else if j > m {
		a.Print(m)
		return
	}
	a.Sort(lo, hi, m)
}

func (a *Array) Partition(lo, hi int) int {

	i := lo + 1
	j := hi
	v := a.arr[lo]

	for {

		for a.arr[i] < v {
			i++
			if i == hi {
				break
			}
		}

		for a.arr[j] > v {
			j--
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		a.Exchange(j, i)
	}
	a.Exchange(j, lo)
	return j
}

func (a *Array) Exchange(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

func (a *Array) Print(j int) {
	for i := 0; i < j; i++ {
		fmt.Print(a.arr[i])
	}
}
