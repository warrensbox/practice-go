package main

/*
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 */
type BinaryMatrix struct {
	Get        func(int, int) int
	Dimensions func() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {

	dim := binaryMatrix.Dimensions()
	rows := dim[0]
	cols := dim[1]
	r := 0
	c := cols - 1
	for r < rows && c >= 0 {
		if binaryMatrix.Get(r, c) == 0 {
			r++
		} else {
			c--
		}
	}

	if c == cols-1 {
		return -1
	}

	return c + 1
}
