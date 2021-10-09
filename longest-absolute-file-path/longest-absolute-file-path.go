package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"

	fmt.Println(lengthLongestPath(input))
}

func lengthLongestPath(input string) int {

	//extension := filepath.Ext("/tmp/hello.go")

	path := strings.Split(input, "\n")
	//fmt.Println("pathHH", path)
	max_so_far := 0
	path_len := 0
	pathHash := make(map[int]int)
	for _, elem := range path {

		level := strings.LastIndexAny(elem, "\t") + 1 //gets the level
		fmt.Println("level", level)
		if level == 0 {
			elem = strings.Replace(elem, "\t", "", -1) //remove the "/t" from the len
			path_len = len(elem)
			fmt.Println("path_len", path_len)
		} else {
			//path[level-1] - get previous path length
			//remove the "/t" from the len using Replace - we do not want 2 tabs to be counted as 2
			// later add back "/" by adding 1
			fmt.Println("pathHash[level-1]", pathHash[level-1])

			elem = strings.Replace(elem, "\t", "", -1)
			fmt.Println("elem", elem)
			path_len = pathHash[level-1] + len(elem) + 1
			fmt.Println("len(elem)", len(elem))
			fmt.Println("path_len", path_len)
		}

		if strings.Contains(elem, ".") { //this is a file
			max_so_far = Max(max_so_far, path_len)
			fmt.Println("max_so_far", max_so_far)
		} else { //not a file (path)
			pathHash[level] = path_len
		}
		fmt.Println(pathHash)
		fmt.Println()
	}
	return max_so_far
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
