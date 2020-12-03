package main

import (
	"fmt"
	"hash/crc32"
)

//define Node
type Node struct {
	Key   string
	Value int
	Next  *Node
}

type List struct {
	head *Node
}

type st struct {
	str [97]List
}

func main() {

	st := st{}
	st.put("hello", 3)
	st.put("hello", 4)
	fmt.Println(st.get("hello"))
	st.put("world", 3)
	fmt.Println(st.get("world"))
	st.put("world", 78)
	fmt.Println(st.get("world"))

}

func (s *st) put(key string, val int) {

	i := hash(key)
	x := s.str[i]
	node := x.head
	for node != nil {
		if key == node.Key {
			node.Value = val
			return
		}
		node = node.Next
	}
	var newNode Node
	newNode.Key = key
	newNode.Value = val

	x.head = &newNode
	s.str[i] = x
}

func (s *st) get(key string) int {

	i := hash(key)
	x := s.str[i]
	node := x.head
	for node != nil {
		if key == node.Key {
			return node.Value
		}
		node = node.Next
	}

	return -1

}
func hash(key string) int {

	return (String(key) & 0x7fffffff) % 97
}

func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
