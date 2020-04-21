package main

import "fmt"

func main() {

	S := "heeellooo"
	words := []string{"hello", "hi", "helo"}

	output := expressiveWords(S, words)

	fmt.Println(output)
}

func expressiveWords(S string, words []string) int {

	res := 0
	for _, val := range words {
		if isSrretchy(S, val) {
			res++
		}
	}
	return res
}

func isSrretchy(s string, w string) bool {
	i := 0
	j := 0

	fmt.Println("s", s)
	fmt.Println("w", w)
	for i < len(s) && j < len(w) {
		fmt.Println("----------")
		c1 := s[i]
		c2 := w[j]

		fmt.Println("c1", string(c1))
		fmt.Println("c2", string(c2))

		if c1 != c2 {
			fmt.Println("First false")
			return false
		}

		e1 := i
		e2 := j

		fmt.Println("e1", e1)
		fmt.Println("e2", e2)

		for e1 < len(s) && s[e1] == c1 {
			e1++
		}
		for e2 < len(w) && w[e2] == c2 {
			e2++
		}

		fmt.Println("e1", e1)
		fmt.Println("e2", e2)

		n1 := e1 - i
		n2 := e2 - j

		fmt.Println("n2", n2)
		fmt.Println("n1", n1)

		if n2 > n1 || (n1 < 3 && n1 != n2) {
			fmt.Println("Second false")
			return false
		}

		i = e1
		j = e2
		if i == len(s) && j == len(w) {
			fmt.Println("Third True")
			return true
		}

	}
	return false
}
