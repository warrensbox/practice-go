package main

import "fmt"

func main() {

}

func checkDuplicates(array []int) {
	bs := BitSet(32000)
	for i := 0; i < len(array); i++ {
		num := array[i]
		num0 := num - 1 //bitset starts at 0, numbers starts at 1
		if bs.get(num0) {
			fmt.Println(num)
		} else {
			bs.set(num0)
		}
	}
}

//BS - set arr
type BS struct {
	bitset []int
}

//BitSet - set BS
func BitSet(size int) *BS {
	bs := BS{}
	bs.bitset = make([]int, (size>>5)+1)
	return &bs
}

func (bs *BS) get(pos int) bool {
	wordNumber := (pos >> 5) + 1
	bitNumber := (pos & 0x1F)
	return (bs.bitset[wordNumber] & (1 << bitNumber)) != 0
}

func (bs *BS) set(pos int) {
	wordNumber := (pos >> 5) + 1
	bitNumber := (pos & 0x1F)
	bs.bitset[wordNumber] |= 1 << bitNumber
}
