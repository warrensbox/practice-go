package main

func maxLength(arr []string) int {

	res := []string{""}
	max := 0
	for _, val := range arr {

		lengthRes := len(res)

		for i := 0; i < lengthRes; i++ {

			newRes := res[i] + val

			restHash := make(map[rune]bool)
			for _, c := range newRes {
				restHash[c] = true
			}
			if len(restHash) != len(newRes) {
				continue
			}

			res = append(res, newRes)
			max = Max(max, len(newRes))

		}

	}

	return max
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
