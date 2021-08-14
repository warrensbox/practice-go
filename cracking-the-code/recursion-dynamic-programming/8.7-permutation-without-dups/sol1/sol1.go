package main

import "fmt"

func main() {

	str := "aacd"
	sample := []rune(str)
	permute(sample, 0, len(sample)-1)
}

/* Function to print permutations of string
This function takes three parameters:
1. String
2. Starting index of the string
3. Ending index of the string. */
func permute(sample []rune, left, right int) {

	//fmt.Println("after l", left)
	if left == right {
		fmt.Println(string(sample))
		return
	}
	hash := make(map[rune]bool)
	for i := left; i <= right; i++ {
		// fmt.Println("l++", l)
		// fmt.Println("i++", i)
		// fmt.Println("l", string(sample[l]))
		//fmt.Println("i", i)
		//fmt.Println("left", left)
		if hash[sample[i]] {
			continue
		} else {
			hash[sample[i]] = true
		}
		sample[left], sample[i] = sample[i], sample[left]
		//fmt.Println("sample B4", string(sample))
		permute(sample, left+1, right)
		// fmt.Println("swap back")
		// fmt.Println("l--", string(sample[l]))
		// fmt.Println("i--", string(sample[i]))
		sample[left], sample[i] = sample[i], sample[left]
		//fmt.Println("sample FtR", string(sample))
	}
}

// /* Function to swap values at two pointers */
// func swap(strx *byte, y byte) {
// 	x, y = y, x
// }
