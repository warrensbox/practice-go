package main

import "fmt"

import "math"

func main() {

	//Input:
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	N := 4
	K := 2

	output := networkDelayTime(times, N, K)

	fmt.Println(output)

}

func networkDelayTime(times [][]int, N int, K int) int {

	dist := make([]int, N+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[K] = 0

	fmt.Println(dist)
	for i := 0; i < N; i++ {
		fmt.Println("----")
		for _, time := range times {
			u := time[0]
			v := time[1]
			w := time[2]

			fmt.Println("u", u)
			fmt.Println("v", v)
			fmt.Println("w", w)
			fmt.Println("dist[u]", dist[u])
			fmt.Println("dist[v]", dist[v])
			if dist[u] != math.MaxInt32 {
				if dist[v] > dist[u]+w {
					dist[v] = dist[u] + w
				}
			}
			fmt.Println("dist[v]", dist[v])
		}
	}

	maxWait := 0

	for i := 1; i <= N; i++ {
		maxWait = int(math.Max(float64(maxWait), float64(dist[i])))
	}

	if maxWait != math.MaxInt32 {
		return maxWait
	} else {
		return -1
	}
}
