package main

import (
	"fmt"
	"math"
)

//https://www.youtube.com/watch?v=Ej9R8SxP9c0

func main() {

	// workers := [][]int{{0, 0}, {1, 1}, {2, 0}}
	// bikes := [][]int{{1, 0}, {2, 2}, {2, 1}}

	workers := [][]int{{0, 0}, {2, 1}}
	bikes := [][]int{{1, 2}, {3, 3}}

	output := assignBikes(workers, bikes)

	fmt.Println(output)
}

func assignBikes(workers [][]int, bikes [][]int) []int {

	workerLength := len(workers)
	bikesLength := len(bikes)
	// result := [workerLength]int{}
	result := make([]int, workerLength)
	assigned := [2000]bool{}
	occupied := [2000]bool{}
	listMap := make(map[int][][]int)

	for i := 0; i < workerLength; i++ {
		for j := 0; j < bikesLength; j++ {
			dist := dist(workers[i], bikes[j])

			fmt.Println("dist", dist)

			//fmt.Println(listMap[dist])

			//dist := getDistance(val[0], val[1])
			//listMap[dist] = append(listMap[dist], dist)
			listMap[dist] = append(listMap[dist], []int{i, j})
		}
	}

	fmt.Println(listMap)

	for i := 0; i < 2001; i++ {
		if _, ok := listMap[i]; ok {

			size := len(listMap[i])
			for j := 0; j < size; j++ {
				arr := listMap[i]
				worker := arr[j][0]
				bike := arr[j][1]
				fmt.Println("wo", worker)
				fmt.Println("bi", bike)

				if !assigned[worker] && !occupied[bike] {
					result[worker] = bike
					assigned[worker] = true
					occupied[bike] = true
				}
			}
		}
	}

	return result
}

func dist(worker []int, bike []int) int {

	// fmt.Println(worker)
	// fmt.Println(bike)

	w1 := math.Abs(float64(worker[0]) - float64(bike[0]))

	w2 := math.Abs(float64(worker[1]) - float64(bike[1]))

	return int(w1 + w2)

}
