package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := []string{"bac", "god", "oxox", "abc", "dog", "cba", "xoxo"}

	hash := make(map[string][]string)
	arr := []string{}
	for _, each := range str {
		key := SortString(each)
		hash[key] = append(hash[key], each)
	}

	for _, val := range hash {
		arr = append(arr, val...)
	}

	fmt.Println(arr)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
