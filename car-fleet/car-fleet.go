package main

import (
	"fmt"
	"sort"
)

func main() {

	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Println(carFleet(target, position, speed))
}

func carFleet(target int, position []int, speed []int) int {

	var allcar Allcar
	for i := 0; i < len(position); i++ {
		time := float64(target-position[i]) / float64(speed[i])
		//fmt.Println(target-position[i])
		//    fmt.Println(speed[i])
		var car Car
		car.Position = position[i]
		car.Time = time
		allcar = append(allcar, car)
	}

	//fmt.Println(allcar)
	sort.Sort(Allcar(allcar))

	fmt.Println(allcar)
	curr := 0.0
	fleet := 0
	//for i:= len(allcar)-1; i >=0; i--{
	for i := 0; i < len(allcar); i++ {
		fmt.Println(allcar[i].Time)
		if allcar[i].Time > curr {
			fmt.Println("fleet", fleet)
			fleet++
			curr = allcar[i].Time
		}
	}
	// fmt.Println(fleet)
	return fleet
}

func (car Allcar) Len() int           { return len(car) }
func (car Allcar) Swap(i, j int)      { car[i], car[j] = car[j], car[i] }
func (car Allcar) Less(i, j int) bool { return car[i].Position > car[j].Position }

type Allcar []Car

type Car struct {
	Position int
	Time     float64
}
