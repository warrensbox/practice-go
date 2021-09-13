package main

import "fmt"

func numSplits(s string) int {

	charIndexMap := make([]int, 26)
	intervals := make([][]int, 0)
	l := len(s)

	//fmt.Println("charIndexMap",charIndexMap)
	//fmt.Println("intervals",intervals)

	for i := 0; i < 26; i++ {
		charIndexMap[i] = -1
	}

	for i := 0; i < l; i++ {
		fmt.Println(string(s[i]))
		slot := s[i] - 'a'
		index := charIndexMap[slot]
		fmt.Println(index)
		if index < 0 {
			//fmt.Println(charIndexMap[slot])
			charIndexMap[slot] = len(intervals)
			intervals = append(intervals, []int{i, i})
			fmt.Println("intervals", intervals)
		} else {
			intervals[index][1] = i
			fmt.Println("intervals", intervals)
		}
	}

	leftCount := 0
	rightCount := len(intervals)
	goodCount := 0

	for i := 0; i < l-1; i++ {
		fmt.Println(string(s[i]))
		index := charIndexMap[s[i]-'a']
		fmt.Println("index", index)
		fmt.Println("i", i)

		if i == intervals[index][0] {
			leftCount++
		}
		if i == intervals[index][1] {
			rightCount--
		}
		fmt.Println("leftCount", leftCount)
		fmt.Println("rightCount", rightCount)
		if leftCount == rightCount {
			goodCount++
		}
	}

	//fmt.Println("charIndexMap",charIndexMap)
	//     for i:=0; i <len(s)-1;i++{
	//         p := s[:i+1]
	//         fmt.Println("p",string(p))

	//         q := s[i+1:]
	//         fmt.Println("q",string(q))
	//     } number-of-good-ways-to-split-a-string

	return 1
}
