package main

func numberOfProvinces(isConnected [][]int) int {

	count := 0
	//spare..
	seen := make(map[int]bool)
	//core
	for i := 0; i < len(isConnected); i++ {
		//parent := i // 2
		if seen[i] {
			continue
		}

		queue := []int{}
		queue = append(queue, i) //2
		//seen[0] seen[1]
		for len(queue) > 0 {

			city := (queue)[0] //qu
			queue = (queue)[1:]
			if seen[city] {
				continue
			}
			seen[city] = true

			for j, present := range isConnected[city] { // 0:1 1:1 2:0
				if !seen[j] {
					if present == 1 {
						seen[j] = true
						queue = append(queue, j)
					}
				}
			}
		}
		count++
	}

	//count 1
	//output
	return count

}
