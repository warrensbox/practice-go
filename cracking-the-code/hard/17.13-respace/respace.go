package main

import (
	"fmt"
	"math"
)

func main() {
	dictionary := make(map[string]bool)
	dictionary["hello"] = true
	dictionary["there"] = true
	dictionary["charley"] = true
	dictionary["cat"] = true
	dictionary["cats"] = true
	dictionary["and"] = true
	dictionary["sand"] = true
	dictionary["dog"] = true
	test := split(dictionary, "hellotherecharley", 0)
	//test := split(dictionary, "catsanddog", 0)
	fmt.Println(test.parsed)

}

func split(dictionary map[string]bool, sentence string, start int) *ParseResult {

	if start >= len(sentence) {
		parseResult := ParseResult{}
		return &parseResult
	}

	bestInvalid := math.MaxInt16
	var bestParsing string
	partial := ""
	index := start

	for index < len(sentence) {
		c := sentence[index]
		partial = partial + string(c)
		//fmt.Println("partial", partial)
		var invalid int
		if _, ok := dictionary[partial]; ok {
			invalid = 0
		} else {
			invalid = len(partial)
		}

		result := split(dictionary, sentence, index+1)
		if invalid+result.invalid < bestInvalid {
			bestInvalid = invalid + result.invalid
			bestParsing = partial + " " + result.parsed
			if bestInvalid == 0 {
				break
			}
		}
		index++
	}

	parseResult := ParseResult{}
	parseResult.invalid = bestInvalid
	parseResult.parsed = bestParsing
	return &parseResult

}

type ParseResult struct {
	invalid int
	parsed  string
}
