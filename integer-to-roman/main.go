package main

func intToRoman(num int) string {

	char := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	mappings := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	curr := ""
	for num > 0 {
		for i := 12; i >= 0; i-- {
			if mappings[i] <= num {
				q := num / mappings[i]
				num = num % mappings[i]

				for k := 0; k < q; k++ {
					curr += char[i]
				}
			}
		}
	}

	return curr
}

/*

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.


12 //i=2

q = 12/10 =1
num = 12%10=2
0-<1
curr = X

2 //i=0

q:= 2/1 = 2
num = 2%1 = 0

0-<2 [0 1]
curr = X+I+I
*/
