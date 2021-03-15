package main

import "fmt"

func main() {

	arr := []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}
	subarr := []int{1, 5, 9}
	getShortest(arr, subarr)
}

func getShortest(arr, subarr []int) {

	//hash := make(map[int][]int)
	hashi := make(map[int]int)

	for _, val := range subarr {
		// hash[val] = []int{}
		hashi[val] = -1000
	}
	minimum := 100000
	minA := len(arr) - 1
	maxA := 0
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("val[%v] index[%v]\n", arr[i], i)
		fmt.Println("____________")
		if _, ok := hashi[arr[i]]; ok {
			//hash[val] = append(hash[val], i)
			hashi[arr[i]] = i
			//	count++
			max, min := checkHash(hashi)
			check := max - min
			fmt.Println("check", check)
			minimum = Min(check, minimum)
			fmt.Println("minimum", minimum)
			if check <= minimum {
				minimum = check
				minA = min
				maxA = max
			}

		}
		// if count == 3 {
		// 	count = 0
		// }
	}
	fmt.Println("hashi", minimum)
	fmt.Println("min", minA)
	fmt.Println("max", maxA)
	fmt.Println(hashi)
}

func checkHash(hashi map[int]int) (int, int) {

	min := 100000
	max := 0
	for k, v := range hashi {
		fmt.Printf("key[%v] value[%v]\n", k, v)
		if v == -1000 {
			return 10000, 1
		}
		min = Min(min, v)
		max = Max(max, v)

	}
	fmt.Printf("MIn[%v] Max[%v]\n", min, max)
	fmt.Println("===DONE===")

	return max, min

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
