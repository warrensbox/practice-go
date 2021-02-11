package main

import "fmt"

func main() {

	str := "abc"
	sample := []rune(str)
	permute(sample, 0, len(sample)-1)
}

/* Function to print permutations of string
This function takes three parameters:
1. String
2. Starting index of the string
3. Ending index of the string. */
func permute(sample []rune, l, r int) {

	fmt.Println("permute")
	if l == r {
		fmt.Println(string(sample))
		return
	}
	for i := l; i <= r; i++ {
		fmt.Println("l++", l)
		fmt.Println("i++", i)
		fmt.Println("l", string(sample[l]))
		fmt.Println("i", string(sample[i]))
		sample[l], sample[i] = sample[i], sample[l]
		permute(sample, l+1, r)
		fmt.Println("swap back")
		fmt.Println("l--", string(sample[l]))
		fmt.Println("i--", string(sample[i]))
		sample[l], sample[i] = sample[i], sample[l]
	}
}

// /* Function to swap values at two pointers */
// func swap(strx *byte, y byte) {
// 	x, y = y, x
// }
