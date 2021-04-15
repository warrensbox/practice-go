package main

import (
	"fmt"
)

func main() {

	s := "cbaebabacd"
	p := "abc"

	output := findAnagrams(s, p)

	fmt.Println(output)
}

func findAnagrams(s string, p string) []int {

	ans := []int{}
	hash := [26]int{}  //s
	phash := [26]int{} //p

	windowofP := len(p)
	lenOfS := len(s)

	for lenOfS < windowofP {
		return ans
	}
	left, right := 0, 0
	for right < windowofP {
		phash[p[right]-'a']++ //p
		hash[s[right]-'a']++  //s
		right++
	}
	right-- // bring it back (increased it previously)

	fmt.Println("right", right)
	// fmt.Println("hash", hash)
	for right < lenOfS { //looping through s

		if phash == hash {
			ans = append(ans, left)
		}
		right++ //move right
		if right != lenOfS {
			hash[s[right]-'a']++
		}
		hash[s[left]-'a']--
		left++
	}

	return ans
}
