package main

func exist(board [][]byte, word string) bool {

	exist := false

	var backtrack func(r, c int, curr string)
	backtrack = func(r, c int, curr string) {

		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] == '1' {
			if curr == word {
				exist = true
			}
			return
		}

		tmp := board[r][c]
		board[r][c] = '1'
		curr += string(tmp)
		backtrack(r+1, c, curr)
		backtrack(r-1, c, curr)
		backtrack(r, c+1, curr)
		backtrack(r, c-1, curr)

		board[r][c] = tmp
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board)[0]; c++ {
			backtrack(r, c, "")
		}
	}

	return exist
}
