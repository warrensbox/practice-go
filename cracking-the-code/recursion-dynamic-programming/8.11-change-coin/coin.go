//https: //www.youtube.com/watch?v=xCbYmUPvc2Q

package main

import "fmt"

func main() {

	total := 5
	coins := 5
	fmt.Println(changeCoin(coins, total))
}

func changeCoin(coins, total int) int {

	//initialize row and column
	db := make([][]int, coins+1)
	for i := range db {
		db[i] = make([]int, total+1)
	}

	for i := 0; i <= coins; i++ {

		for j := 0; j <= total; j++ {

			if i == 0 && j == 0 {
				db[i][j] = 1
			} else {
				if i > j {
					db[i][j] = db[i-1][j]
				} else if i >= 1 {
					db[i][j] = db[i-1][j] + db[i][j-i]
				}
			}
		}
	}

	// for i := 0; i <= coins; i++ {
	// 	for j := 0; j <= total; j++ {
	// 		fmt.Print(db[i][j])
	// 	}
	// 	fmt.Println()
	// }

	return db[coins][total]
}
