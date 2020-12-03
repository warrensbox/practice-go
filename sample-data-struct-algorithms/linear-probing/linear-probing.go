package main

import (
	"fmt"
	"hash/crc32"
)

const M = 300011

//define Node
type arr struct {
	Keys   [M]string
	Values [M]int
}

func main() {

	pb := arr{}
	pb.put("hello", 1)
	fmt.Println(pb.get("hello"))

	pb.put("hello", 7)
	fmt.Println(pb.get("hello"))
}

func (pb *arr) put(key string, val int) {

	i := hash(key)
	for pb.Keys[i] != "" {
		if pb.Keys[i] == key {
			break
		}
		i = (i + 1) % M
	}

	pb.Keys[i] = key
	pb.Values[i] = val

}

func (pb *arr) get(key string) int {

	i := hash(key)
	for pb.Keys[i] != "" {
		if pb.Keys[i] == key {
			return pb.Values[i]
		}
		i = (i + 1) % M
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
