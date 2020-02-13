package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "11"
	b := "1"

	ans := addBinary(a, b)

	fmt.Println(ans)
}

func addBinary(a string, b string) string {

	n := len(a)
	m := len(b)

	if n < m {
		return addBinary(b, a)
	}

	L := max(n, m)

	// result.WriteString(
	var result strings.Builder
	carry := 0
	j := m - 1

	for i := L - 1; i > -1; i-- {

		if string(a[i]) == "1" {
			carry++
		}

		if j > -1 {
			//need to minus j --
			if string(b[j]) == "1" {
				carry++
			}
		}

		if carry%2 == 1 {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}

		carry = carry / 2

		if carry == 1 {
			result.WriteString("1")
		}
	}

	//TODO reverse string and return
	return result.String()
}

// class Solution {
// 	public String addBinary(String a, String b) {
// 	  int n = a.length(), m = b.length();
// 	  if (n < m) return addBinary(b, a);
// 	  int L = Math.max(n, m);

// 	  StringBuilder sb = new StringBuilder();
// 	  int carry = 0, j = m - 1;
// 	  for(int i = L - 1; i > -1; --i) {
// 		if (a.charAt(i) == '1') ++carry;
// 		if (j > -1 && b.charAt(j--) == '1') ++carry;

// 		if (carry % 2 == 1) sb.append('1');
// 		else sb.append('0');

// 		carry /= 2;
// 	  }
// 	  if (carry == 1) sb.append('1');
// 	  sb.reverse();

// 	  return sb.toString();
// 	}
//   }

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
