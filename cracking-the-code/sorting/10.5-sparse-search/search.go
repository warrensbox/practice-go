package main

import "fmt"

func main() {

	arr := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	fmt.Println(search(arr, "dad", 0, len(arr)-1))
}

func search(arr []string, str string, first, last int) int {

	if first > last {
		return -1
	}

	//move mid to the middle
	mid := (last + first) / 2

	//if mid is empty, find closest non-empty string
	if arr[mid] == "" {
		left := mid - 1
		right := mid + 1
		for true {
			if left <= first && right > last {
				return -1
			} else if right <= last && arr[right] != "" {
				mid = right
				break
			} else if left >= first && arr[right] != "" {
				mid = left
				break
			}
			right++
			left--
		}
	}

	//check for string, and recurse if necessary
	if arr[mid] == str { //found it
		return mid
	} else if str > arr[mid] {
		return search(arr, str, mid+1, last)
	} else { //if str < arr[mid]
		return search(arr, str, first, mid-1)
	}
}
