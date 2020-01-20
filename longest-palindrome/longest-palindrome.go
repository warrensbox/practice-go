package main

import "fmt"

func main(){

	input := "ashannahmom"
	output := longestPalindrome(input)

	fmt.Println(output)
}


func longestPalindrome(s string) string {
	
	max := ""
	for i := 0; i < len(s);i++ {
		fmt.Println("-----------LONG ----------- ")
		s1:=extend(s,i,i) //case for mom
		s2:=extend(s,i,i+1) //case for hannah


		fmt.Println("s1",s1)
		fmt.Println("s2",s2)
		

 		if ( len(s1) > len(max)) {
			 max = s1
		 }
 		if ( len(s2) > len(max)) {
			 max = s2
		 }

		 fmt.Println(" ----------- LONG ----------- ")
	}

    return max
}


func extend(s string, i int,j int) string {

	fmt.Println("######## EXTEND ###########")

	for i >= 0 && j < len(s){

		fmt.Println("i",i)
		fmt.Println("j",j)
		
		fmt.Println("s-i",string(s[i]))
		fmt.Println("s-j",string(s[j]))

		if string(s[i]) != string(s[j]){
			break
		}
		
		i--
		j++ 

	}

	fmt.Println("i ---",i)
	fmt.Println("j ---",j)
	fmt.Println("substring",string(s[i+1:j]))
	fmt.Println("######## EXTEND ###########")
	return string(s[i+1:j])

}