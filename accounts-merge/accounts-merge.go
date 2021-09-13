package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {

	edges := make(map[int][]int)
	emailToIdx := make(map[string]int)
	for i, acc := range accounts {
		for j := 1; j < len(acc); j++ {
			if idx, exists := emailToIdx[acc[j]]; exists {
				edges[idx] = append(edges[idx], i)
				// fmt.Println("1",edges)
				edges[i] = append(edges[i], idx)
				// fmt.Println("2",edges)
			} else {
				emailToIdx[acc[j]] = i
			}
		}
	}

	//fmt.Println(emailToIdx)
	var res [][]string
	visited := make(map[int]bool)
	for i, acc := range accounts {
		if visited[i] {
			continue
		}
		// fmt.Println("IIIiii",i)
		//bfs
		queue := []int{i}
		acctEmails := make(map[string]bool)
		for len(queue) > 0 {
			idx := queue[0]
			queue = queue[1:]
			if !visited[idx] {
				visited[idx] = true
				for _, email := range accounts[idx][1:] {
					acctEmails[email] = true
				}
				//fmt.Println(acctEmails)
				for _, neighbor := range edges[idx] {
					//fmt.Println("neighbor",neighbor)
					queue = append(queue, neighbor)
				}
			}
			// fmt.Println("here")
		}
		//  fmt.Println("map",acctEmails)
		mergedAcc := []string{acc[0]}
		for email := range acctEmails {
			mergedAcc = append(mergedAcc, email)
		}

		sort.Strings(mergedAcc[1:])
		res = append(res, mergedAcc)
	}

	return res
}
