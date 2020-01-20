package main

import "fmt"

func main(){

	strArr := []string{"flower","flow","flight"}
	ans := longestCommonPrefix(strArr)

	fmt.Println(ans)
}


func longestCommonPrefix(strs []string) string {
	
	longest := ""

	if len(strs) == 0 {
		return longest
	}


	comparisonWord := strs[0]
	fmt.Println("comparisonWord",comparisonWord)
	comparisonIndex := 0
	fmt.Println("comparisonIndex",comparisonIndex)

	for _, comparisonLetter :=  range comparisonWord {
		fmt.Println("comparisonLetter", string(comparisonLetter))

		for i:=1; i< len(strs); i++ {

			currentWord := strs[i]
			fmt.Println("currentWord",currentWord)
			currentLetter := string(currentWord[comparisonIndex])
			fmt.Println("currentLetter",currentLetter)


			if (string(comparisonLetter) != currentLetter || comparisonIndex > len(currentWord)){
				return longest;
			}
		}

		comparisonIndex =comparisonIndex+1
		fmt.Println("comparisonIndex",comparisonIndex)
		longest = longest+string(comparisonLetter)
		fmt.Println("------------------------")
	}

	return longest
}