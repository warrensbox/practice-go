package main

func checkStraightLine(coordinates [][]int) bool {

	//     sort.Slice(coordinates, func(a, b int) bool{
	//         return coordinates[a][0] < coordinates[b][0]
	//     })

	if len(coordinates) <= 2 {
		return true
	}
	dX := coordinates[1][0] - coordinates[0][0]
	dY := coordinates[1][1] - coordinates[0][1]
	for i := 2; i < len(coordinates); i++ {

		rX := coordinates[i][0] - coordinates[0][0]
		rY := coordinates[i][1] - coordinates[0][1]

		if dY*rX != dX*rY {
			return false
		}
	}

	return true
}
