package main

func removeDuplicateLetters(s string) string {

	// this lets us keep track of what's in our solution in O(1) time
	seen := make(map[rune]bool)
	// this will let us know if there are any more instances of s[i] left in s
	lastSeen := make(map[rune]int)

	for i, char := range s {
		lastSeen[char] = i
	}

	stack := []rune{}

	for i := 0; i < len(s); i++ {
		val := rune(s[i])
		// we can only try to add c if it's not already in our solution
		// this is to maintain only one of each character
		if _, ok := seen[val]; !ok { //if not seen

			for len(stack) > 0 && stack[len(stack)-1] > val && i < lastSeen[stack[len(stack)-1]] {
				delete(seen, stack[len(stack)-1])
				stack = stack[:len(stack)-1]

			}

			seen[val] = true
			stack = append(stack, val)
		}
	}

	return string(stack)
}
