package main

func tictactoe(moves [][]int) string {

	n := 3
	rows := make([]int, n)
	cols := make([]int, n)
	diag := 0
	anti := 0
	player := 1
	for _, move := range moves {
		row := move[0]
		col := move[1]

		rows[row] += player
		cols[col] += player

		if row == col {
			diag += player
		}

		if row+col == n-1 {
			anti += player
		}

		if Abs(rows[row]) == n || Abs(cols[col]) == n || Abs(diag) == n || Abs(anti) == n {
			return playerID(player)
		}

		player *= -1
	}

	if len(moves) == n*n {
		return "Draw"
	}
	return "Pending"
}

func playerID(a int) string {
	if a == 1 {
		return "A"
	} else {
		return "B"
	}
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
