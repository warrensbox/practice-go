package main

import (
	"fmt"
	"math"
)

func main() {

	//Input:
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	N := 4
	K := 2

	output := networkDelayTime(times, N, K)

	fmt.Println(output)

}

type node struct {
	dst int
	wt  int
}

type qnode struct {
	vertex int
	dist   int
}

func networkDelayTime(times [][]int, N int, K int) int {

	adj := make(map[int][]*node)

	for _, packet := range times {
		source := packet[0]
		dest := packet[1]
		wt := packet[2]
		adj[source] = append(adj[source], &node{dest, wt})
	}

	queue := []*qnode{}

	queue = append(queue, &qnode{K, 0})

	// for key, val := range adj {
	// 	fmt.Println("key", key)
	// 	fmt.Println(val)
	// }

	distance := make([]int, N+1)
	for i := range distance {
		distance[i] = math.MaxInt32
	}

	distance[K] = 0
	for len(queue) > 0 {
		//top queue
		curr := (queue)[0]
		queue = (queue)[1:]
		fmt.Println("queue", queue)
		//currDist := curr.wt
		fmt.Println("top", curr)
		for _, t := range adj[curr.vertex] {
			tmp := qnode{}
			tmp.vertex = t.dst
			tmp.dist = t.wt
			fmt.Println("t", t)
			if distance[tmp.vertex] > tmp.dist {
				queue = append(queue, &tmp)
				distance[tmp.vertex] = t.wt + curr.dist
			}
		}
	}

	// for _, val := range distance {
	// 	fmt.Println("val", val)
	// }

	maxWait := 0

	for i := 1; i <= N; i++ {
		maxWait = int(math.Max(float64(maxWait), float64(distance[i])))
	}

	if maxWait != math.MaxInt32 {
		return maxWait
	} else {
		return -1
	}
}
