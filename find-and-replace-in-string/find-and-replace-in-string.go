package main

import (
	"fmt"
	"sort"
)

func main() {

	S := "abcd"
	indexes := []int{0, 2}
	sources := []string{"a", "cd"}
	targets := []string{"eee", "ffff"}

	output := findReplaceString(S, indexes, sources, targets)

	fmt.Println(output)
}

type bundle struct {
	index  int
	source string
	target string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {

	bundles := make([]bundle, len(indexes))

	for i := 0; i < len(indexes); i++ {
		bundles[i] = bundle{
			index:  indexes[i],
			source: sources[i],
			target: targets[i],
		}
	}

	fmt.Println("string", S)
	sort.Slice(bundles, func(i, j int) bool {
		return bundles[i].index > bundles[j].index
	})

	fmt.Println(bundles)
	fmt.Println()
	for i := 0; i < len(bundles); i++ {
		bundle := bundles[i]

		fmt.Println("len(S)", len(S))
		fmt.Println("bundle.index", bundle.index)
		fmt.Println("len(bundles)", len(bundles))
		fmt.Println(S[bundle.index : bundle.index+len(bundle.source)])
		fmt.Println("bundle.source", bundle.source)

		if len(S) < bundle.index+len(bundle.source) || S[bundle.index:bundle.index+len(bundle.source)] != bundle.source {
			fmt.Println("continue")
			continue
		}
		fmt.Println("S[:bundle.index]", S[:bundle.index])
		fmt.Println("S[bundle.index+len(bundle.source):]", S[bundle.index+len(bundle.source):])

		S = S[:bundle.index] + bundle.target + S[bundle.index+len(bundle.source):]
		fmt.Println("S", S)
		fmt.Println()
	}
	return S
}
