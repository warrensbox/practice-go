package main

import (
	"fmt"
)

type Constructor struct {
	arr []int
}

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	construct := &Constructor{
		arr: arr,
	}
	//fmt.Println(&construct)

	profile := Initialize(construct)

	profile.union(4, 3)
	profile.union(3, 8)
	profile.connected(1, 8)
	fmt.Println(profile.arr)

}

func Initialize(attr *Constructor) *Constructor {

	for i := 0; i < len(attr.arr); i++ {
		attr.arr[i] = i
	}

	return attr
}

func (attr *Constructor) connected(p, q int) {

	if attr.arr[p] == attr.arr[q] {
		fmt.Println("true")
		return
	}

	fmt.Println("false")
}

func (attr *Constructor) union(p, q int) {

	pid := attr.arr[p]
	qid := attr.arr[q]

	for i := 0; i < len(attr.arr); i++ {
		if attr.arr[i] == pid {
			attr.arr[i] = qid
		}
	}

}
