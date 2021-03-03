package main

import "fmt"

func main() {

	fmt.Println(factorial(10))
	fmt.Println(countFactZeros(10))
}

func factorial(x int) int {
	//base case
	if x == 1 {
		return x
	}

	return x * factorial(x-1)
}

func factorsOf5(i int) int {
	//fmt.Println("i", i)
	count := 0
	for i%5 == 0 {
		//fmt.Println("5", i)
		count++
		i /= 5
	}
	//fmt.Println("COUNT", count)
	return count
}

func countFactZeros(num int) int {
	count := 0
	for i := 2; i <= num; i++ {
		count += factorsOf5(i)
		//fmt.Println("KOUNT", count)
	}
	return count
}
