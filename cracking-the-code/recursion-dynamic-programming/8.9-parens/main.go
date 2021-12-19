package main

import "fmt"

func main() {
	Paren(3)
}

var arr []string

func Paren(parenthesis int) {
	paren("", parenthesis, parenthesis)
	fmt.Println(arr)
}

func paren(str string, left, right int) {

	if left <= right && left >= 0 {
		fmt.Println("left", left)
		fmt.Println("right", right)
		if left == 0 && right == 0 {
			fmt.Println(str)
			arr = append(arr, str)
			fmt.Println()
		}
		paren(str+"(", left-1, right)
		paren(str+")", left, right-1)
	}

}
