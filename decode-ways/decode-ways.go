package main

//watch https://www.youtube.com/watch?v=YcJTyrG3bZs
import (
	"fmt"
	"strconv"
)

func main() {

	input := "226"
	output := numDecodings(input)
	fmt.Println(output)

}

func numDecodings(s string) int {

	n := len(s)

	if n == 0 {
		return 0
	}

	memo := make([]int, n+1)

	memo[n] = 1

	fmt.Println(memo)
	fmt.Println("s", s)

	ints := []rune(s)[n-1]
	fmt.Println("ints", ints)
	fmt.Println("n", n)
	if ints != '0' {
		memo[n-1] = 1
	} else {
		memo[n-1] = 0
	}
	fmt.Println("memo2", memo)
	for i := n - 2; i >= 0; i-- {
		intd := []rune(s)[i]
		if intd == '0' {
			continue
		} else {

			mem, _ := strconv.Atoi(s[i : i+2])

			if mem <= 26 {
				memo[i] = memo[i+1] + memo[i+2]
			} else {
				memo[i] = memo[i+1]
			}

		}
	}

	return memo[0]
}

/* Explanation */
/*
Assigning memo[n] = 1; means when the string is empty, there is only one answer.
memo[n-1] = s.charAt(n-1) != '0' ? 1 : 0; means when there is only one character in the string,
if this character is not 0, there will be an answer, or there will be no answer.
Then it starts the dp portion. When we add a letter from the end of the string, if the first two letters <=26, we can get memo[n]=memo[n+1]+memo[n+2].
For example, the String now is "123xxxx" and we know all the result from 2.
Because 12<26, we can make this string either"12"+"3xxxx" or 1+23xxxx which is exactly memo[n]=memo[n-1]+memo[n-2].
if the String is"32xxxx" memo[n]=memo[n+1]. if there are 0s in the string, we should skip it and look at the next character because there is no answer when the string begins with 0.
*/
