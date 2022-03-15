package main

import "fmt"

/*

keep two pointers
i = 0
j = len(arr)-1

point to the front and point to the back

if arr[i] > arr [j] {
	arr[i+1] += arr[i]
	count++
	i++
}else if arr[j] > arr [j]{
	arr[j-1] += arr[j]
	count++
	j--
}else{
	i++
	j--
}

*/

func main() {

	arr := []int{1, 6, 4, 1, 3, 7}
	fmt.Println(CountNum(arr))
}

func CountNum(arr []int) int {

	count := 0
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i] < arr[j] {
			arr[i+1] += arr[i]
			count++
			i++
		} else if arr[j] < arr[i] {
			arr[j-1] += arr[j]
			count++
			j--
		} else {
			i++
			j--
		}
	}

	return count
}
