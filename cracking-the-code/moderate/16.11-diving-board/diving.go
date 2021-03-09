package main

import "fmt"

func main() {

	fmt.Println(DivingBoard(3, 1, 2))
}

func DivingBoard(k int, shorter, longer int) map[int]int {

	length := make(map[int]int)
	divingBoardRecusive(k, 0, shorter, longer, length)
	return length
}

func divingBoardRecusive(k int, total, shorter, longer int, length map[int]int) {
	if k == 0 {
		length[total] = 1
		return
	}

	divingBoardRecusive(k-1, total+shorter, shorter, longer, length)
	divingBoardRecusive(k-1, total+longer, shorter, longer, length)
}
