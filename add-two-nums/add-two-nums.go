package main

import "fmt"


func main(){

	arr1 := []int{2,4,3}
	arr2 := []int{5,6,4}

	ans := addTwoNumbers(arr1,arr2)

	fmt.Println(ans)
}


 func addTwoNumbers(l1 []int, l2 []int) []int {
 
	a,b,c,x,y,z := 0,0,0,0,0,0

	//inverse the values
	//add numbers together as a multiple of 10 each time
	for i := len(l1); i > 0 ; i-- {

		x = l1[i-1]
		y = x+z
		z = y*10

		a = l2[i-1]
		b = a+c
		c = b*10

	}

	newSum := y+b
	remains := 0
 
	var arr []int
	

	for newSum > 0 {
		remains = newSum%10
		 arr = append(arr,remains)
		 newSum = newSum/10
	}

	return arr
}