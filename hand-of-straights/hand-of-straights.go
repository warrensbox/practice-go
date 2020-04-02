package main

import (
	"fmt"
	"sort"
)

func main() {

	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	W := 3

	output := isNStraightHand(hand, W)

	fmt.Println(output)

}

type IntCount struct {
	value int
	count int
}

func isNStraightHand(hand []int, W int) bool {

	if W == 1 {
		return true
	}

	if len(hand)%W != 0 {
		return false
	}

	dict := make(map[int]int)

	for _, val := range hand {
		dict[val]++
	}

	fmt.Println(dict)

	counts := []IntCount{}

	for k, v := range dict {
		counts = append(counts, IntCount{
			value: k,
			count: v,
		})
	}

	//fmt.Println(counts)

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].value < counts[j].value
	})

	//fmt.Println(counts)

	n := len(hand) / W

	for n > 0 {
		fmt.Println("*****************")
		counts[len(counts)-1].count--
		for i := 1; i < W; i++ {
			// fmt.Println("I=", i)
			// fmt.Println("len(counts) - 1 - i =", len(counts)-1-i)
			// fmt.Println("counts[len(counts)-1-i].value + 1 =", counts[len(counts)-1-i].value+1)
			// fmt.Println("counts[len(counts)-i].value = ", counts[len(counts)-i].value)

			if len(counts)-1-i < 0 || counts[len(counts)-1-i].value+1 != counts[len(counts)-i].value {
				return false
			}

			counts[len(counts)-1-i].count--
			//fmt.Println(counts)
		}

		p := len(counts) - W

		// fmt.Println("p ", p)
		// fmt.Println(counts)

		for i := len(counts) - W; i < len(counts); i++ {
			//fmt.Println("i ", i)
			//fmt.Println("counts[i].count ", counts[i].count)

			if counts[i].count != 0 {
				// fmt.Println("counts[p] ", counts[p])
				// fmt.Println("counts[i] ", counts[i])

				counts[p] = counts[i]
				p++
			}
		}
		counts = counts[:p]
		//fmt.Println("COUNTS", counts)

		n--
	}
	return true
}

///ANOTHER SOLUTION
/*
func isNStraightHand(hand []int, W int) bool {
    if len(hand)%W != 0 {
        return false
    }
    sort.Ints(hand)
    hash := make(map[int]int)
    for _, num := range hand {
        hash[num]+= 1
    }

    for _, num := range hand {
        if hash[num] == 0 {
            continue
        }
        hash[num]-=1
        if !dfs(hash, W-1, num+1)  {
            //fmt.Println(hash, num)
            return false
        }
    }
    return true
}

func dfs(hash map[int]int, window, num int) bool {
    if window == 0 {
        return true
    }
    if hash[num] == 0 {
        return false
    }
    hash[num]-=1
    return dfs(hash, window-1, num+1)
}

*/
