package main

func main() {

}
func exist(board [][]byte, word string) bool {

	exist := false

	var backtrack func(r, c int, wordBack string) bool

	backtrack = func(r, c int, wordBack string) bool {
		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] == '1' {
			return false
		}

		if board[r][c] != wordBack[0] {
			return false
		}

		if len(wordBack) == 1 {
			return true
		}

		tmp := board[r][c]
		board[r][c] = '1'

		if backtrack(r+1, c, wordBack[1:]) {
			return true
		}
		if backtrack(r-1, c, wordBack[1:]) {
			return true
		}
		if backtrack(r, c+1, wordBack[1:]) {
			return true
		}
		if backtrack(r, c-1, wordBack[1:]) {
			return true
		}

		board[r][c] = tmp

		return false
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if backtrack(r, c, word) {
				return true
			}
		}
	}

	return exist
}
