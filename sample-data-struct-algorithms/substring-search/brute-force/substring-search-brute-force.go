package main

import "fmt"

func main() {

	pat := "tomato"

	txt := "peoplelovetomatoes"

	fmt.Println(search(pat, txt))
}

func search(pat, txt string) bool {

	m := len(pat) //length of pattern
	n := len(txt) //length of text

	for i := 0; i <= n-m; i++ {
		var j int
		for j = 0; j < m; j++ { //iterate though the pattern
			if txt[i+j:i+j+1] != pat[j:j+1] {
				break
			}
		}
		if j == m {
			return true
		}
	}

	return false
}
