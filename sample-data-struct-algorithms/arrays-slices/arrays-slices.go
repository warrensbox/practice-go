package main

import "fmt"

func main(){

	arr := [6]string{"to", "be", "or", "not","to","be"}
					//0     1      2     3    4     5

	fmt.Println(arr[:6]) //[to be or not to be] (6-0) //index 0-5

	fmt.Println(arr[5:]) //[be]  (6-5=1) //index 5-end

	fmt.Println(arr[:4]) //[to be or not] (4-0) //index 0-3

	fmt.Println(arr[2:4]) //[or not] // (4-2 = 2) //index 2-3

	fmt.Println(arr[1:]) //[be or not to be] // (6-1 = 5) //index 1-end

}