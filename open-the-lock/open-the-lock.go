package main

import (
	"fmt"
	"strconv"
)

func main() {

	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"

	output := openLock(deadends, target)

	fmt.Println(output)
}

func openLock(deadends []string, target string) int {

	res := -1

	visited := make(map[string]bool)

	for _, val := range deadends {
		visited[val] = true
	}

	if visited["0000"] {
		return res
	}

	visited["0000"] = true
	queue := []string{"0000"}
	rotate := 0

	for len(queue) > 0 {
		newQueue := []string{}
		for _, node := range queue {
			fmt.Println("node", node)
			if node == target {
				return rotate
			}
			fmt.Println("len(node)", len(node))
			for i := 0; i < len(node); i++ {

				fmt.Println("---START---")
				fmt.Println("i", i)

				fmt.Println("int(node[i]-'0')", int(node[i]-'0'))
				j := (int(node[i]-'0') + 1) % 10
				fmt.Println("j", j)
				fmt.Println("node[:1]", node[:1])
				fmt.Println("node[i+1:]", node[i+1:])
				ns := node[:i] + strconv.Itoa(j) + node[i+1:]

				fmt.Println("ns", ns)

				if !visited[ns] {
					visited[ns] = true
					newQueue = append(newQueue, ns)
				}

				fmt.Println("------")

				j = (int(node[i]-'0') - 1 + 10) % 10
				fmt.Println("j", j)
				fmt.Println("node[:1]", node[:1])
				fmt.Println("node[i+1:]", node[i+1:])
				ns = node[:i] + strconv.Itoa(j) + node[i+1:]
				fmt.Println("ns", ns)
				if !visited[ns] {
					visited[ns] = true
					newQueue = append(newQueue, ns)
				}

			}
		}

		// break
		queue = newQueue
		fmt.Println("QUEUE", queue)
		rotate++

	}
	return res
}

//try with channels

func openLock2(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, dead := range deadends {
		visited[dead] = true
	}

	if visited["0000"] {
		return -1
	}
	visited["0000"] = true
	queue := []string{"0000"}
	rotate := 0
	for len(queue) > 0 {
		newQueue := []string{}
		for _, s := range queue {
			if s == target {
				return rotate
			}
			fmt.Println("s", s)
			for i := 0; i < len(s); i++ {
				j := (int(s[i]-'0') + 1) % 10
				ns := s[:i] + strconv.Itoa(j) + s[i+1:]
				fmt.Println("node[:1]", s[:1])
				fmt.Println("node[i+1:]", s[i+1:])
				fmt.Println("ns", ns)
				if !visited[ns] {
					visited[ns] = true
					newQueue = append(newQueue, ns)
				}
				j = (int(s[i]-'0') - 1 + 10) % 10
				ns = s[:i] + strconv.Itoa(j) + s[i+1:]
				fmt.Println("ns", ns)
				if !visited[ns] {
					visited[ns] = true
					newQueue = append(newQueue, ns)
				}
			}
		}
		if rotate == 1 {
			break
		}

		queue = newQueue
		rotate++
	}
	return -1
}
