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
}

func solution(service_list []string, entrypoint string) []string {

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
		node := (queue)[0]
		(queue) = (queue)[1:]
		mapSum[node]++
		for _, val := range graph[node] {
			queue = append(queue, val)
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
