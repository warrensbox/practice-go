package main

import "fmt"

func main() {
	arr := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	num := Array{}
	num.arr = arr
	num.Sort(0, len(arr)-1)
	fmt.Println(num.arr)
}

type Array struct {
	arr []int
}

func (a *Array) Sort(lo, hi int) {

	if hi <= lo {
		return
	}

	j := a.Partition(lo, hi)
	a.Sort(lo, j-1)
	a.Sort(j+1, hi)
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
