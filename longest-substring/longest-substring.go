package main

import "fmt"

func main(){

	//ans := lengthOfLongestSubstring("abcabcs")
	ans := lengthOfLongestSubstringBruteForce("abcabcs")

	fmt.Println(ans)
}

func lengthOfLongestSubstringBruteForce(s string) int {
	length := len(s)
	max := 0
	for i := 0; i < length; i++{
		for j := i+1; j <= length; j++{
			num := checkUnique(s,i,j)
			if num > max {
				max =num
			}

		}
	}
	fmt.Println("Final",max)
	return max  
}


func checkUnique(s string, start int, end int) int{

	dict := make(map[string]int)
	counter := 0
	for i := start; i < end; i++{
		_, exist := dict[string(s[i])]
		if exist{
			//check if character is already in dictionary
			return counter
		}else{
			dict[string(s[i])] = 1
			counter++
		}
	}

	return counter
}

// public int lengthOfLongestSubstring(String s) {
// 	int n = s.length();
// 	Set<Character> set = new HashSet<>();
// 	int ans = 0, i = 0, j = 0;
// 	while (i < n && j < n) {
// 		// try to extend the range [i, j]
// 		if (!set.contains(s.charAt(j))){
// 			set.add(s.charAt(j++));
// 			ans = Math.max(ans, j - i);
// 		}
// 		else {
// 			set.remove(s.charAt(i++));
// 		}
// 	}
// 	return ans;
// }



//func lengthOfLongestSubstring(s string) int {

// 	lengthOfString := len(s)
// 	dict := make(map[string]int)
// 	ans,i,j := 0,0,0
// 	for ( i < lengthOfString && j < lengthOfString ){
// 		fmt.Println(i)
// 		fmt.Println(j)
// 		 _,exist := dict[string(s[j])]
// 		 if !exist {
// 			fmt.Println("enter")
// 			 j = j+1
// 			 dict[string(s[j])] = 1
// 			 ans = max(ans, j - i)
// 		 }else{
// 			i = i+1
// 			delete(dict, string(s[i]));
// 		 }
// 	}


// 	//fmt.Println(ans)
//  return 1
// }



func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

