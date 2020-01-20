package main


import (
	"fmt"
	"strconv"
)


func main(){

	input := -12321
 	output := isPalindrome(input)

	fmt.Println(output)
}

func isPalindrome(x int) bool {

	str:= strconv.Itoa(x)
	output := isPalindromeString(str)

	output2 := isPalindromeInt(x)

	fmt.Println(output2)
	
	return output
}


func isPalindromeString(str string) bool{

	
	i :=0 
	j := len(str)-1
	for i < len(str)/2 {

		if (string(str[i]) !=  string(str[j])) {
			fmt.Println("Not")
			return false
		}
		i++
		j--
	}

	return true
} 

func isPalindromeInt(x int) bool{


	inverseVal := inverseNumber(x)

	add := inverseVal+x
	multi := inverseVal*2
	fmt.Println(add)

	fmt.Println(multi)

	if add == multi {
		fmt.Println("Valid")
		return true
	}else{
		fmt.Println("InValid")
		return false
	}

}

func inverseNumber(x int) int{


	newInt := 0
	value :=0
	negative := false


	if x <0 {
		return 0
	}

	for x > 0 {

		remains := x%10
		value = remains+newInt
		newInt = value*10
		x = x/10
	}

	if negative {
		value = value*-1
	}


	return value

}