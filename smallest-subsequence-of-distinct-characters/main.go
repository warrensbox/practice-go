package main

func smallestSubsequence(s string) string {

	lastSeen := make(map[rune]int)
	seen := make(map[rune]bool)
	for i, char := range s {
		lastSeen[char] = i
	}

	stack := []rune{}

	for i, val := range s {

		if _, ok := seen[val]; !ok {

			for len(stack) > 0 && val < stack[len(stack)-1] && lastSeen[stack[len(stack)-1]] > i {
				delete(seen, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, val)
			seen[val] = true
		}

	}

	return string(stack)
}
