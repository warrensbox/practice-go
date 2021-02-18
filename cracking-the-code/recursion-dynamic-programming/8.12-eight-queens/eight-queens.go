package main

import "fmt"

func main() {

	columns := make([]int, GRID)
	results := [][]int{}
	placeQueen(0, columns, results)
}

//var results []int

const (
	GRID = 4
)

func placeQueen(row int, columms []int, results [][]int) {
	if row == GRID {
		results = append(results, columms)
		fmt.Println(results)
	} else {
		for col := 0; col < GRID; col++ {
			if checkValid(columms, row, col) {
				columms[row] = col //place queen
				placeQueen(row+1, columms, results)
			}
		}
	}
}

/*
  check if (row1,coumn1)is a valid spot for a queen by checking if there is a
  queen in the same column or diagonal. We don't need to check it for queens in
  the same row because the calling placeQueen only attempts to place on queen at
  a time. We this row is empty
*/
func checkValid(columns []int, row1 int, column1 int) bool {

	for row2 := 0; row2 < row1; row2++ {
		column2 := columns[row2]
		/*
			Check if (row2,cloumn2) incalidates (row1,column1) as a queen spot
			Check if rows have a queen in the same colum
		*/
		if column1 == column2 {
			return false
		}

		/*
			Check  diagonals: if the distance between the columns equals the distance
			between the rows, then they are in the same diagonal
		*/
		columnDistance := Abs(column2 - column1)

		//row1 > row2, so no need for abs
		rowDistance := row1 - row2
		if columnDistance == rowDistance {
			return false
		}
	}
	return true
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
