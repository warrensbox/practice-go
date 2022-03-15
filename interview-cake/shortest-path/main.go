package main

import (
	"fmt"
)

/*

 Goal : find shortest path to node (A->B)
 Sol :
 - use bfs
 - mark seen nodes (keep track)
 - take from and dest names
 O(N+M).

 The tricky part was backtracking to assemble the path we used to reach our endNode. In general, it's helpful to think of backtracking as two steps:
 howWeReachedNodes
*/

func main() {

	hashGraph := make(map[string][]string)
	hashGraph["Min"] = []string{"William", "Jayden", "Omar"}
	hashGraph["William"] = []string{"Min", "Noam"}
	hashGraph["Jayden"] = []string{"Min", "Amelia", "Ren", "Noam"}
	hashGraph["Ren"] = []string{"Jayden", "Omar"}
	hashGraph["Amelia"] = []string{"Jayden", "Adam", "Miguel"}
	hashGraph["Adam"] = []string{"Amelia", "Miguel", "Sofia", "Lucas"}
	hashGraph["Miguel"] = []string{"Amelia", "Adam", "Liam", "Nathan"}
	hashGraph["Noam"] = []string{"Nathan", "Jayden", "William"}
	hashGraph["Omar"] = []string{"Ren", "Min", "Scott"}

	shortestPath(hashGraph, "Min", "Liam")

}

func shortestPath(hashGraph map[string][]string, from, dest string) {

	if from == dest {
		return
	}
	var queue []string
	queue = append(queue, from)
	hasSeen := make(map[string]bool)
	howWeReachedNodes := make(map[string]string)
	hasSeen[from] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := (queue)[0]
			queue = (queue)[1:]
			connections := hashGraph[node]

			if dest == node {
				findParent(howWeReachedNodes, dest)
				return
			}

			for _, val := range connections {

				if !hasSeen[val] {
					hasSeen[val] = true
					queue = append(queue, val)
					howWeReachedNodes[val] = node
				}

			}

		}
	}

}

func findParent(hash map[string]string, child string) {

	currentNode := child
	var path []string
	for currentNode != "" {
		path = append(path, currentNode)
		currentNode = hash[currentNode]
	}

	fmt.Println(path)
}
