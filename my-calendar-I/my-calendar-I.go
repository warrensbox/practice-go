package main

import "fmt"

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func main() {
	obj := Constructor()
	param := obj.Book(3, 5)
	param = obj.Book(1, 4)
	fmt.Println(param)
}

type MyCalendar struct {
	Calendar [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {

	fmt.Println("this.Calendar", this.Calendar)

	for i := 0; i < len(this.Calendar); i++ {
		if start < this.Calendar[i][1] && end > this.Calendar[i][0] {
			return false
		}
	}

	this.Calendar = append(this.Calendar, []int{start, end})
	return true
}
