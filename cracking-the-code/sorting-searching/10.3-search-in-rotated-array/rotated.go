package main

import "fmt"

func main() {

	arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	//arr, element, left, right
	fmt.Println(search(arr, 5, 0, len(arr)-1))
}

func search(arr []int, element int, left, right int) int {

	mid := (left + right) / 2
	if element == arr[mid] { //base case
		return mid
	}
	if right < left { //base case
		return -1
	}

	if arr[left] < arr[mid] { //left side is normally ordered
		if element >= arr[left] && element < arr[mid] {
			return search(arr, element, left, mid-1) //search left
		} else {
			return search(arr, element, mid+1, right) //search right
		}
	} else if arr[mid] < arr[left] { // right is normally ordered
		if element > arr[mid] && element <= arr[right] {
			return search(arr, element, mid+1, right) //search right
		} else {
			return search(arr, element, left, mid-1) //search left
		}
	} else if arr[left] == arr[mid] {
		if arr[mid] != arr[right] { //if right is different, search it
			return search(arr, element, mid+1, right) //search right
		} else { //we have to search both sides
			result := search(arr, element, left, mid-1) //search left
			if result == -1 {
				return search(arr, element, mid+1, right) //search right
			} else {
				return result
			}
		}
	}

	return -1
}
