package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {

	hashMap := make(map[string][]string)
	badList := []string{}
	for i := range keyName {
		hashMap[keyName[i]] = append(hashMap[keyName[i]], keyTime[i])
	}

	fmt.Println(hashMap)

	for key, val := range hashMap {
		times := val
		sort.Strings(times)

		for i := 2; i < len(times); i++ {
			curr := getTime(times[i])
			last := getTime(times[i-2])
			if curr-last <= 60 {
				badList = append(badList, key)
				break
			}

		}
	}

	sort.Strings(badList)

	return badList
}

func getTime(time string) int {

	hourMin := strings.Split(time, ":")
	hour, _ := strconv.Atoi(hourMin[0])
	min, _ := strconv.Atoi(hourMin[1])
	min += hour * 60

	return min
}

/*

save everything in a hash map

dan. | 10, 1040, 11 ,12
louis | 9,11, 13, 15

step sort array of each map (dan)

loop through the sorted array starting at index 2 (look back -2)

*/
