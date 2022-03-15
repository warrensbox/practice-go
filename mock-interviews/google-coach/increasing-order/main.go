package main

import "fmt"

/*

start =1
end =1
max = Max(max, start-end+1)

1,2,4,5, 3

ends-map
end:start
3:1
5:4


starts-map
start:end
1:2
4:5

psuedo
when nothing on hash
set both start and end to val

when is in hash -> get the key from the starts-map
set start to key value
end to val

- check ends-map
- if val-1 => map1,
	if val + 1 > map2 { combine it}
	yes - get the val of map1 (set it to start), delete key of start, put (val,start)
	no - add to map1 (set start and end to val)


- if val+1 => map, yes - , no - add to map (set start and end to val)
	yes - get the val of map (set it to end), delete key of start, put (val,end)
*/

func main() {

	seq := &Seq{
		make(map[int]int),
		make(map[int]int),
		0,
	}
	seq.Add(1)
	seq.Add(2)
	seq.Longest()

	seq.Add(4)
	seq.Add(10)
	seq.Add(11)
	seq.Longest()
	seq.Add(3)
	seq.Longest()
	seq.Add(9)
	seq.Add(8)
	seq.Add(6)
	seq.Add(7)
	seq.Longest()
}

type Seq struct {
	Ends     map[int]int
	Starts   map[int]int
	maxSoFar int
}

func (s *Seq) Add(val int) {

	start, startExist := s.Ends[val-1]
	end, endExist := s.Starts[val+1]

	if startExist && endExist { //if both val-1 (in ends map) and val+1 (in starts map) exist, combine it (STICK 'EM TOGETHER)
		delete(s.Ends, val-1)
		delete(s.Ends, end)
		delete(s.Starts, val+1)
		delete(s.Starts, start)
		s.Ends[end] = start
		s.Starts[start] = end
		s.maxSoFar = Max(s.maxSoFar, end-start+1)

	} else if startExist { //if only val-1 exist (STICK IT TO THE BACK)
		s.Ends[val] = start   //create the start value [ends map]
		delete(s.Ends, val-1) // remove the existing entry in the [ends map]
		s.Starts[start] = val // update the entry [starts map]
		s.maxSoFar = Max(s.maxSoFar, val-start+1)

	} else if endExist { //if only val+1 exist (STICK IT TO THE FRONT)
		s.Starts[val] = end     //create the end value [starts map]
		delete(s.Starts, val+1) // remove the existing entry in the [starts map]
		s.Ends[end] = val       // update the entry [ends map]
		s.maxSoFar = Max(s.maxSoFar, end-val+1)
	} else {
		s.Ends[val] = val //put entry in both
		s.Starts[val] = val
		s.maxSoFar = Max(s.maxSoFar, 1)
	}

	fmt.Println("ends", s.Ends)
	fmt.Println("starts", s.Starts)
}

func (s *Seq) Longest() int {
	fmt.Println(s.maxSoFar)
	return s.maxSoFar
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
