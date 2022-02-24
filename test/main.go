package main

import "fmt"

/*

 Design a time-based key-value data structure that can store multiple values for
 the same key at different timestamps and retrieve the key's value at a certain
 timestamp.

*/

func main() {

	keyval := New()
	keyval.Set("foo", "bar", 1)
	fmt.Println(keyval.Get("foo", 1))
	keyval.Set("cat", "dog", 2)
	keyval.Set("day", "night", 3)
	fmt.Println(keyval.Get("day", 4))
	keyval.Set("foo", "fighters", 4)
	keyval.Set("foo", "apple", 6)
	fmt.Println(keyval.Get("foo", 5))

}

type TimeMap struct {
	hashMap map[string][]Values
}

type Values struct {
	time int
	val  string
}

// Returns a new TimeMap
func New() *TimeMap {
	return &TimeMap{
		hashMap: make(map[string][]Values),
	}
}

// Stores the key with the corresponding value at the given timestamp.
func (t *TimeMap) Set(key, value string, timestamp int) {

	var properties Values
	properties.val = value
	properties.time = timestamp
	//check key exist; if exist => append it to existing key
	//if not, create new key - val
	t.hashMap[key] = append(t.hashMap[key], properties)
}

// Retrieves the value associated with the given key at a timestamp such as
// `key.timestamp <= timestamp`. If there are multiple values, return the value
// with the largest `key.timestamp`. If there are no values, return an empty
// string.
func (t *TimeMap) Get(key string, timestamp int) string {
	res := ""
	//check if key exist
	properties, ok := t.hashMap[key]
	if !ok {
		return res
	}
	//properties {1,barr} {4, fighter} {5, apple}

	if len(properties) == 1 {
		return properties[0].val
	}

	left := 0
	right := len(properties) - 1
	target := timestamp
	for left < right {

		mid := (left + right) / 2

		if target < properties[mid].time {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return properties[left].val
}
