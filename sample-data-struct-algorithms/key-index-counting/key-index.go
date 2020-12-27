package main

import "fmt"

func main() {

	letters := []rune{'d', 'a', 'c', 'f', 'f', 'b', 'd', 'b', 'f', 'b', 'e', 'a'}

	aux := make([]rune, len(letters))
	count := make([]int, 7)

	for i := 0; i < len(letters); i++ {
		count[(letters[i]-'a')+1]++ //count occurances +1
	}

	for i := 0; i < 6; i++ {
		count[i+1] += count[i] //compute cummulative
	}

	for i := 0; i < len(letters); i++ {
		val := letters[i]       //get value of the element
		index := count[val-'a'] //lookup index of the element in count array
		aux[index] = val        //set value to aux array based on the index obtained from above
		count[val-'a']++        //increase count after
	}

	copy(letters, aux) //copy aux to letters

	//print to verify
	for i := 0; i < len(letters); i++ {
		fmt.Print(string(letters[i]))
	}

}
