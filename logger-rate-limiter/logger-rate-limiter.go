package main

import (
	"fmt"
)

func main() {

	obj := Constructor()

	param_1 := obj.ShouldPrintMessage(1, "foo")
	param_2 := obj.ShouldPrintMessage(2, "bar")
	param_3 := obj.ShouldPrintMessage(3, "foo")
	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
}

type Logger struct {
	dict map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{dict: map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {

	time, ok := this.dict[message]
	if ok && timestamp-time < 10 {
		return false
	}

	this.dict[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
