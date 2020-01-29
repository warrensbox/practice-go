package main

import (
	"fmt"
)

func main() {

	input := "23"
	output := threeSum(input)
	fmt.Println(output)
}

func threeSum(digits string) []string {

	var output []string
	var arr []string
	if len(digits) != 0 {
		backtrack("", digits, 0, arr)
	}

	return output
}

// private void combination(String prefix, String digits, int offset, List<String> ret) {
// 	if (offset >= digits.length()) {
// 		ret.add(prefix);
// 		return;
// 	}
// 	String letters = KEYS[(digits.charAt(offset) - '0')];
// 	for (int i = 0; i < letters.length(); i++) {
// 		combination(prefix + letters.charAt(i), digits, offset + 1, ret);
// 	}
// }

func backtrack(prefix string, digit string, offset int, arr []string) {

	if offset >= len(digit) {
		arr = append(arr, prefix)
	}
	dict := make(map[string]string)
	dict["2"] = "abc"
	dict["3"] = "def"
	dict["4"] = "ghi"
	dict["5"] = "jkl"
	dict["6"] = "mno"
	dict["7"] = "pqrs"
	dict["8"] = "tuv"
	dict["9"] = "wxyz"

	letters := dict[string(digit[offset])]

	fmt.Println("letters", letters)
	fmt.Println("string(digit[offset]", string(digit[offset]))

	for i := 0; i < len(letters); i++ {
		fmt.Println(i)
		letter := string(letters[i])
		fmt.Println(letter)
		//fmt.Println(digit)
		// 		combination(prefix + letters.charAt(i), digits, offset + 1, ret);
		backtrack(prefix+letter, digit, offset+1, arr)
	}

}
