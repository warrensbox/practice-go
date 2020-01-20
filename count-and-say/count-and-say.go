package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){

	input := 5
	output := countAndSay(input)

	fmt.Println(output)
}

func countAndSay(n int) string {

	var result strings.Builder

	sequence1 := "1"
	//if n is less than or = 1, return 1
	if n <= 1 {
		return sequence1
	}
	//else
	temp := sequence1
	var j int
	// i starts from 2 because we know the first sequnce already
	for i:=2;i<=n;i++{
		fmt.Println("------outside-------")
		fmt.Println("i", i)
		fmt.Println("len(temp)", len(temp))
		for j=0;j<len(temp);j++{
			fmt.Println("#####inner######")
			fmt.Println("j", j)
			count := 1
			fmt.Println("count", count)
			tempStr := temp[j]
			fmt.Println("tempStr", string(tempStr))

			for j < len(temp)-1 && temp[j]==temp[j+1]{
				fmt.Println("-----inner-inner--------")
				count++
				j++
				fmt.Println("count", count)
				fmt.Println("j", j)
				fmt.Println("-----inner-inner--------")
			}
			
			result.WriteString(strconv.Itoa(count)+string(tempStr))
			fmt.Println("result", strconv.Itoa(count)+string(tempStr))
			fmt.Println("#####inner######")
		}
		temp =  result.String()
		fmt.Println("temp", temp)
		result.Reset()
		fmt.Println("------outside-------")
	}
	return temp
}