package main

func findFinalValue(nums []int, original int) int {

	hash := make(map[int]int)

	for _, val := range nums {
		hash[val] = val * 2
	}

	for {

		if newVal, ok := hash[original]; ok {
			original = newVal
		} else {
			break
		}
	}

	return original
}
