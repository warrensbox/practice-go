package main

import (
	"fmt"
	"sort"
)

const maxLen = 10

func main() {
	queries := []string{"bbbbf", "cc"}
	words := []string{"cccccc", "ccddd", "ccc", "cccc"}

	output := numSmallerByFrequency(queries, words)

	fmt.Println(output)
}

// func numSmallerByFrequency(queries []string, words []string) []int {

// 	fr := make([]int, maxLen+2)

// 	for _, w := range words {

// 		fr[f(w)]++
// 		fmt.Println("fr", fr)
// 	}

// 	return fr
// }

// func f(s string) int {
// 	fmt.Println("----")
// 	arr := make([]int, 26)
// 	min := 25
// 	for _, c := range s {

// 		i := int(c - 'a')

// 		arr[i]++

// 		if i < min {
// 			min = i
// 		}
// 	}
// 	fmt.Println("min", min)
// 	fmt.Println("arr[min]", arr[min])

// 	return arr[min]

// }

func numSmallerByFrequency(queries []string, words []string) []int {
	var ans []int
	count := make([]int, len(words))
	for i, w := range words {
		fmt.Println("f(w)", f(w))
		count[i] = f(w)
	}
	fmt.Println("count", count)
	sort.Ints(count)
	fmt.Println("count", count)
	for _, q := range queries {
		fmt.Println("queries", q)
		fmt.Println("f(q)+1", f(q)+1)
		fmt.Println("sort.SearchInts(count, f(q)+1)", sort.SearchInts(count, f(q)+1))
		ans = append(ans, len(count)-sort.SearchInts(count, f(q)+1))
	}
	return ans
}
func f(s string) int {
	n := make([]int, 26)
	for i := 0; i < len(s); i++ {
		n[s[i]-'a']++
	}
	for i := 0; i < len(n); i++ {
		if n[i] != 0 {
			return n[i]
		}
	}
	return 0
}
