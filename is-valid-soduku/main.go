package main

func main() {

}

func isValidSudoku(board [][]byte) bool {

	/* check horizontally */
	/* check vertically */
	/* check in 3 x 3 matrix */

	/*
	   0	1,2,3 | 4,5,6 | 7,8,9
	   1	3 4 7 | 1
	   2	4 6 8
	   	_
	   	4
	   	5
	   	6
	   	7
	   	8
	   	9

	   MapHorizontal
	   Key : Value
	   1   : [0 3]
	   2   :
	   3   :
	   4   : [1 1]

	*/

	return check3by3(board) && checkVertical(board) && checkHorizontal(board)

}

func check3by3(board [][]byte) bool { //O(n^2)

	for row := 0; row < len(board); row += 3 { // O(n)
		for col := 0; col < len(board[0]); col += 3 { // O(n)
			mem := make(map[byte]bool)
			for r := row; r < row+3; r++ { // O(3)
				for c := col; c < col+3; c++ { // O(3)
					if mem[board[r][c]] {
						return false
					} else {
						mem[board[r][c]] = true
					}
				}
			}
		}
	}
	return true
}

func checkVertical(board [][]byte) bool { //O(n^2)

	for col := 0; col < len(board[0]); col++ {
		mem := make(map[byte]bool)
		for row := 0; row < len(board); row++ {
			if mem[board[row][col]] {
				return false
			} else {
				mem[board[row][col]] = true
			}
		}
	}

	return true
}

func checkHorizontal(board [][]byte) bool {
	//O(n^2)

	for row := 0; row < len(board); row++ {
		mem := make(map[byte]bool)
		for col := 0; col < len(board[0]); col++ {
			if mem[board[row][col]] {
				return false
			} else {
				mem[board[row][col]] = true
			}
		}
	}

	return true
}
