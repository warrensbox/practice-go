package main

func reverseStr(s string, k int) string {

	str := []rune(s)

	for i := 0; i < len(str); i += 2 * k {

		left := i
		right := Min(i+(k-1), len(str)-1)
		for left < right {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
	}

	return string(str)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
