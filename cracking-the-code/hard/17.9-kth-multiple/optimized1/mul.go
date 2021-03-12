package main

import (
	"container/list"
	"fmt"
)

/*
3 -> 5 -> 7 ->
5 -> 7 -> 9 -> 15 -> 21 ->
7 -> 9 -> 15 -> 21 -> 15 -> 25 -> 35 ->
9 -> 15 -> 21 -> 15 -> 25 -> 35 -> 21 -> 35 -> 49 ->
*/

func main() {
	fmt.Println(getkthMagicNumber(4))
}

func addProducts(q *list.List, v int) {
	q.PushBack(v * 3)
	q.PushBack(v * 5)
	q.PushBack(v * 7)
}

func removeMin(q *list.List) int {

	min := q.Front().Value.(int)
	for list := q.Front(); list != nil; list = list.Next() {
		if min > list.Value.(int) {
			min = list.Value.(int)
		}
	}
	current := q.Front()
	for current != nil {

		alpha := current.Value.(int)
		if alpha == min {
			current = current.Next()
			q.Remove(current.Prev())
		} else {
			current = current.Next()
		}

	}
	return min
}

func getkthMagicNumber(k int) int {
	if k < 0 {
		return 0
	}

	val := 1
	q := list.New()
	addProducts(q, 1)
	for i := 0; i < k; i++ {
		val = removeMin(q)
		addProducts(q, val)
	}
	return val
}

func PrintAll(q *list.List) {
	for list := q.Front(); list != nil; list = list.Next() {
		fmt.Printf("%v -> ", list.Value)
	}
	fmt.Println()
}
