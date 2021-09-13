package main

import (
	"bytes"
	"fmt"
)

func groupStrings(strings []string) [][]string {

	var res [][]string

	hash := make(map[string][]string)
	for _, word := range strings {
		len := len(word)

		if len == 0 {
			hash[word] = append(hash[word], word)
		}

		longerWord := shift(word, word[0]-'a')
		fmt.Println("longerWord", longerWord)
		hash[longerWord] = append(hash[longerWord], word)
		fmt.Println("----")
	}

	for _, v := range hash {
		//fmt.Println("k",k)
		// fmt.Println("v",v)
		res = append(res, v)
	}

	//fmt.Println("----")
	// fmt.Println(hash)

	return res
}

func shift(s string, count byte) string {

	fmt.Println("count", count)
	var sb bytes.Buffer
	for i := 0; i < len(s); i++ {
		fmt.Println("tet", s[i]-'a'+26-count)
		ch := byte((s[i]-'a'+26-count)%26 + 'a')
		sb.WriteByte(ch)
	}
	fmt.Println("str", sb.String())
	return sb.String()
}
