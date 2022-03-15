package main

func findDisappearedNumbers(num []int) []int {

	output := []int{}
	for i := 0; i < len(num); i++ {
		index := Abs(num[i]) - 1

		if num[index] > 0 {
			num[index] = -num[index]
		}

	}

	for i := 0; i < len(num); i++ {
		if num[i] > 0 {
			output = append(output, i+1)
		}
	}

	return output
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*


nums = [4,3,2,7,8,2,3,1]

rearrange = [1,2,3,4,6,7,8]


*/
