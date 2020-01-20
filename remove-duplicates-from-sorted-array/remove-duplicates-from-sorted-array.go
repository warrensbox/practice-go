package main

import "fmt"

func main(){

	inputArr := []int{1,1,2,2,3,4}

	ans := removeDuplicates(inputArr)

	fmt.Println(ans)
}

func removeDuplicates(nums []int) int {


	lengthArr := len(nums)

	if lengthArr ==0 { return 0}
	i := 0
	for j :=1; j < lengthArr; j++{

		if nums[i] != nums[j]{
			i++
			nums[i]=nums[j]
		}

	}


	return i+1
    
}