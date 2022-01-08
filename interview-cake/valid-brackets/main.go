package main

import "fmt"

/*
goal :find if opening and closing brackets match (accordingly to type)
sol :
- {} () []
- ({[]})
- (){[]}

- iterate "({})"
- stack [ ( {   ]
- hash map to memorize bracket types
hash[")"] = (
hash["]"] = []
hash["}"] = {

*/

func main() {

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(validator("{[(])}"))

}

func validator(strBrackets string) bool {

	hash := make(map[rune]rune)
	hash['('] = ')'
	hash['{'] = '}'
	hash['['] = ']'
	var stack []rune
	for _, brac := range strBrackets {

		switch brac {
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			} else {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if brac != hash[top] {
					return false
				}
			}
		default:
			stack = append(stack, brac)
		}

	}

	return len(stack) == 0

}
