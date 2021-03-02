package main

func main() {

}

func hasWon(board [][]string, row, col int) string {

	if len(board) != len(board[0]) {
		return "Empty"
	}
	piece := board[row][col]
	if hasWonRow(board, row) || hasWonCol(board, col) {
		return piece
	}
	if row == col && hasWonDiagonal(board, 1) {
		return piece
	}
	if row == (len(board)-col-1) && hasWonDiagonal(board, -1) {
		return piece
	}
	return "Empty"
}

func hasWonRow(board [][]string, row int) bool {

	for c := 1; c < len(board[row]); c++ {
		if board[row][c] != board[row][0] {
			return false
		}
	}
	return true
}

func hasWonCol(board [][]string, column int) bool {

	for r := 1; r < len(board); r++ {
		if board[r][column] != board[0][column] {
			return false
		}
	}
	return true
}

func hasWonDiagonal(board [][]string, direction int) bool {

	row := 0
	var column int
	if direction == 0 {
		column = 0
	} else {
		column = len(board) - 1
	}
	first := board[0][column]
	for i := 0; i < len(board); i++ {
		if board[row][column] != first {
			return false
		}
		row += 1
		column += direction
	}
	return true
}
