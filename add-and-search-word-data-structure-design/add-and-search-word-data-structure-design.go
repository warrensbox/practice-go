package main

import "fmt"

func main() {

	obj := Constructor()
	fmt.Println(obj)
	obj.AddWord("hello")
	output1 := obj.Search("hello")

	fmt.Println(output1)

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

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	return this.searchHelper(word, 0, this.root)
}

func (this *WordDictionary) searchHelper(word string, pos int, curr *Node) bool {
	if pos >= len(word) {
		return curr.isEnd
	}

	r := word[pos]
	if r == '.' {
		if len(curr.links) == 0 {
			return false
		}

		pos++
		for _, link := range curr.links {
			if this.searchHelper(word, pos, link) {
				return true
			}
		}
	}

	if link, ok := curr.links[rune(r)]; ok {
		pos++
		return this.searchHelper(word, pos, link)
	}

	return false
}
