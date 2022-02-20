package main

import "fmt"

func main() {
	str1 := "anna"
	goalbank := []string{"nana", "dana", "banana", "aann"}
	fmt.Println(swapingArr(str1, goalbank))
}

func swapingArr(str string, goals []string) []bool {

	res := []bool{}

	for _, val := range goals {
		output := canSwap(str, val)
		res = append(res, output)
	}

	return res

}

func canSwap(str, goal string) bool {

	if len(str) != len(goal) {
		return false
	}

	if len(str) < 2 || len(goal) < 2 {
		return false
	}

	if len(str) == 0 || len(goal) == 0 {
		return false
	}
	//memorize map
	hash := make(map[byte]bool)

	for _, val := range goal {
		hash[byte(val)] = true
	}

	counter := 0
	stash := []int{}
	for i := 0; i < len(goal); i++ {

		if counter > 2 {
			return false
		}

		if str[i] != goal[i] {
			counter++
			stash = append(stash, i)
		}

	}

	for _, val := range stash {
		if !hash[str[val]] {
			return false
		}
	}

	return true
}

/*


parent(
    iterate thru goalbank [i ..n ]
        output := each val - > sub func (canswap) returns true or false
        res = aapend(res, output)

    return res
)

str1 = anna
goalbank = [nana , dana, banana, aann]

return [true, false, false, true]

nana
0123

n | true
a | true

counter =2 (0 1)

aNna
0123 (0 1)


1) track index (ssee if you can swap == gaol)

2) len(str) == len(goal) false

3) len(str) == 0 || len(goal) false

4) counter > 2 false

5)

1 map -> goal
2 iterate goal and str - get the index that does not match (str)
3 check hash

time : O(n)
space : O(n)


*/
