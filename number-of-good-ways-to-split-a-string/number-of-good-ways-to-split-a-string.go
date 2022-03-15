package main

import "fmt"

func main() {

	s := "aacaba"
	//112233
	fmt.Println(numSplits(s))
}

func numSplits(s string) int {

	mapLeft := make(map[byte]bool)
	mapRight := make(map[byte]bool)
	prefix := make([]int, len(s))
	suffix := make([]int, len(s))
	prefix[0] = 1
	mapLeft[s[0]] = true
	for i := 1; i < len(s); i++ {
		if mapLeft[s[i]] {
			prefix[i] = prefix[i-1]
		} else {
			mapLeft[s[i]] = true
			prefix[i] = prefix[i-1] + 1
		}
	}

	suffix[len(s)-1] = 1
	mapRight[s[len(s)-1]] = true
	for i := len(s) - 2; i >= 0; i-- {
		if mapRight[s[i]] {
			suffix[i] = suffix[i+1]
		} else {
			mapRight[s[i]] = true
			suffix[i] = suffix[i+1] + 1
		}
	}

	goodCount := 0
	for i := 1; i < len(s); i++ {

		if prefix[i-1] == suffix[i] {
			goodCount++
		}
	}
	return goodCount
}

func numSplits2(s string) int {

	charIndexMap := make([]int, 26)
	intervals := make([][]int, 0)
	l := len(s)

	//fmt.Println("charIndexMap",charIndexMap)
	//fmt.Println("intervals",intervals)

	for i := 0; i < 26; i++ {
		charIndexMap[i] = -1
	}

	for i := 0; i < l; i++ {
		fmt.Println("i", i)
		fmt.Println(string(s[i]))
		slot := s[i] - 'a'
		index := charIndexMap[slot]
		fmt.Println(index)
		if index < 0 {
			fmt.Println("len(intervals)", len(intervals))
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
	fmt.Println("rightCount", rightCount)
	for i := 0; i < l-1; i++ {
		fmt.Println(string(s[i]))
		index := charIndexMap[s[i]-'a']
		fmt.Println("index", index)
		fmt.Println("i", i)
		fmt.Println("intervals[index][0]", intervals[index][0])
		fmt.Println("intervals[index][1]", intervals[index][1])
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
