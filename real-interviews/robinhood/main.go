package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	service_list := []string{"logging=",
		"user=logging",
		"orders=user,foobar",
		"recommendations=user,orders",
		"dashboard=user,orders,recommendations"}

	entrypoint := "dashboard"

	fmt.Println(solution(service_list, entrypoint))

	//    # Output (note sorted by service name)
	//    ["dashboard*1",
	// 	"logging*4",
	// 	"orders*2",
	// 	"recommendations*1",
	// 	"user*4"]

}

func solution(service_list []string, entrypoint string) []string {

	//graph
	//graph := make(map[Node][]Node)
	graph := make(map[string][]string)
	for _, val := range service_list {

		services := strings.Split(val, "=") //https://pkg.go.dev/strings#Split
		serviceName := services[0]          //dashboard
		dependencies := services[1]         //user,orders,recommendations
		arrDep := strings.Split(dependencies, ",")

		for _, child := range arrDep {
			graph[serviceName] = append(graph[serviceName], child)
		}

	}

	queue := []string{entrypoint}
	mapSum := make(map[string]int)

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := (queue)[0]
			(queue) = (queue)[1:]
			child, _ := graph[node]
			mapSum[node]++
			for _, val := range child {
				queue = append(queue, val)
			}

		}

	}

	res := []string{}
	for k, v := range mapSum {
		if k != "non_existent" && k != "" {
			res = append(res, fmt.Sprintf("%s*%v", k, v))
		}

	}

	sort.Strings(res)
	return res

}
