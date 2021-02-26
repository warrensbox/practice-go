package main

import (
	"fmt"
)

func main() {
	l := Listy{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l.arr = append(l.arr, arr...)

	fmt.Println(search(l, 15))
}

type Listy struct {
	arr []int
}

func (l *Listy) elementAt(i int) int {

	if i < len(l.arr)-1 {
		return l.arr[i]
	}
	return -1
}

func search(l Listy, value int) int {

	index := 1
	for l.elementAt(index) != -1 && l.elementAt(index) < value {
		index *= 2
		fmt.Println("index", index)
	}
	fmt.Println("over", index)
	return binarySearch(l, value, index/2, index)
}

func binarySearch(l Listy, value int, lo, hi int) int {
	var mid int
	for lo <= hi {
		mid = (lo + hi) / 2
		midVal := l.elementAt(mid)
		if midVal > value || midVal == -1 {
			hi = mid - 1
		} else if midVal < value {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
