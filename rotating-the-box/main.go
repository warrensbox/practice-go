package main

func rotateTheBox(box [][]byte) [][]byte {

	width := len(box)
	length := len(box[0])

	rotatedwidth := length
	rotatedlength := width

	//init
	rotate := make([][]byte, rotatedwidth)
	for i := range rotate {
		rotate[i] = make([]byte, rotatedlength)
	}

	for row := 0; row < width; row++ {
		empty := length - 1
		for col := length - 1; col >= 0; col-- {
			if box[row][col] == '*' {
				empty = col - 1
			}
			if box[row][col] == '#' {
				box[row][col] = '.'
				box[row][empty] = '#'
				empty--
			}
		}
	}

	for row := 0; row < width; row++ {
		for col := length - 1; col >= 0; col-- {
			rotate[col][width-1-row] = box[row][col]
		}
	}

	return rotate

}
