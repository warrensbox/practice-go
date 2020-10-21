package main

import "fmt"

func main() {

	//listExist := make(map[int]bool)

	obj := Constructor()

	obj.AddWord("hello")

}

type WordDictionary struct {
	root *Node
}

type Node struct {
	isEnd bool
	links map[rune]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			links: make(map[rune]*Node),
		},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for _, r := range word {
		fmt.Println("string(r)", string(r))
		node, ok := curr.links[r]
		fmt.Println("ok", ok)
		if !ok {
			curr.links[r] = &Node{
				links: make(map[rune]*Node),
			}
			curr = curr.links[r]
		} else {
			curr = node
		}
	}
	curr.isEnd = true
}

/** Adds a word into the data structure. */
func (this *WordDictionary) PrintAll() {
	curr := this.root

	fmt.Println()
	curr.isEnd = true
}
