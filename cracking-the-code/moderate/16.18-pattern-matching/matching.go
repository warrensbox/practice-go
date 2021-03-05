package main

import "reflect"

func main() {

}

// func doesMatch(pattern string, value string) bool {

// 	if len(pattern) == 0 {
// 		return len(value) == 0
// 	}

// 	mainChar := pattern[0]
// 	var altChar byte
// 	if mainChar == 'a' {
// 		altChar = 'b'
// 	} else {
// 		altChar = 'a'
// 	}
// 	size := len(value)

// 	countOfMain := countOf(pattern, mainChar)
// 	countOfAlt := len(pattern) - countOfMain
// 	firstAlt := IndexOf(altChar, pattern)
// 	maxMainSize := size / countOfMain

// 	///More to do
// }

func countOf(pattern string, c byte) int {
	count := 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == c {
			count++
		}
	}
	return count
}

func IndexOf(params ...interface{}) int {
	v := reflect.ValueOf(params[0])
	arr := reflect.ValueOf(params[1])

	var t = reflect.TypeOf(params[1]).Kind()

	if t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}
	return -1
}
