package main

import "fmt"

func main() {
	fmt.Println(WaysToGetTrue("1^0", true))
	//fmt.Println(WaysToGetTrue("1^0|0|1", false))
	//fmt.Println(WaysToGetTrue("0&0&0&1^1|0", true))
}

func WaysToGetTrue(str string, boolean bool) int {

	return waysToGetTrue(str, 0, len(str)-1, boolean)
}

func waysToGetTrue(str string, i, j int, isTrue bool) int {

	if i > j {
		return 0
	}

	if i == j {
		if isTrue {
			if str[i] == '1' {
				return 1
			}
		} else {

			if str[i] == '0' {
				return 1
			}
		}
	}
	//fmt.Println("reseting an")
	ans := 0
	for k := i + 1; k < j; k = k + 2 {

		leftTrue := waysToGetTrue(str, i, k-1, true)
		leftFalse := waysToGetTrue(str, i, k-1, false)
		rightTrue := waysToGetTrue(str, k+1, j, true)
		rightFalse := waysToGetTrue(str, k+1, j, false)

		// fmt.Println("ans", ans)
		fmt.Println("leftTrue", leftTrue)
		fmt.Println("leftFalse", leftFalse)
		fmt.Println("rightTrue", rightTrue)
		fmt.Println("rightFalse", rightFalse)
		//fmt.Println("str[k]", str[k])
		if str[k] == '^' {
			if isTrue {
				fmt.Println("^ true")
				ans = ans + leftTrue*rightFalse + leftFalse*rightTrue
			} else {
				ans = ans + leftTrue*rightTrue + leftFalse*rightFalse
			}
		} else if str[k] == '|' {
			if isTrue {
				ans = ans + leftFalse*rightTrue + leftTrue*rightFalse + leftTrue*rightTrue
			} else {
				ans = ans + leftFalse*rightFalse
			}
		} else {
			//this if for &
			if isTrue {
				ans = ans + leftTrue*rightTrue
			} else {
				ans = ans + leftFalse*rightFalse + leftFalse*rightTrue + leftTrue*rightFalse
			}
		}
	}

	return ans
}
