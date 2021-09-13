package main

import "strconv"

func maximumSwap(num int) int {

	hash := make(map[int]int)
	bs := []byte(strconv.Itoa(num))
	for i := 0; i < len(bs); i++ {
		hash[int(bs[i]-'0')] = i
	}

	for i := 0; i < len(bs); i++ {
		for j := 9; j > int(bs[i]-'0'); j-- {

			if val, ok := hash[j]; ok {
				if hash[j] > i {

					bs[i], bs[val] = bs[val], bs[i]

					byteNumber := []byte(bs)
					byteToInt, _ := strconv.Atoi(string(byteNumber))
					return byteToInt
				}
			}
		}
	}

	return num
}
