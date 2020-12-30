package main

//see example https://golangbyexample.com/trie-implementation-in-go/

import (
	"container/list"
	"fmt"
)

func main() {

	trie := NewTrie()
	trie.Put("hello", "world")
	trie.Put("hel", "no")
	trie.Put("vi", "va")
	trie.Put("she", "hotel")
	trie.Put("shells", "hotel")
	// //trie.Put("hey", "there")
	// fmt.Println("-------")
	// fmt.Println(trie.Get("hello"))
	// fmt.Println("-------")
	// fmt.Println(trie.Del("hello"))
	// fmt.Println("-------")
	// fmt.Println(trie.Get("hel"))

	//trie.keys()

	//trie.keysWithPrefix("hel")

	trie.longestPrefix("shells")
}

const (
	//ALBHABET_SIZE total characters in english alphabet
	ALBHABET_SIZE = 26
)

type TrieNode struct {
	val       string
	childrens [ALBHABET_SIZE]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	newTrie := &Trie{
		root: &TrieNode{},
	}
	return newTrie
}

func (t *Trie) Put(key, val string) {
	wordLength := len(key)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := key[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &TrieNode{}
		}
		current = current.childrens[index]
	}
	current.val = val
}

func (t *Trie) Get(key string) string {
	wordLength := len(key)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := key[i] - 'a'
		if current.childrens[index] == nil {
			return ""
		}
		current = current.childrens[index]
	}
	if current.val != "" {
		return current.val
	}
	return ""
}

func (t *Trie) GetNode(key string) *TrieNode {
	wordLength := len(key)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := key[i] - 'a'
		if current.childrens[index] == nil {
			return nil
		}
		current = current.childrens[index]
	}

	return current
}

func (t *Trie) Del(key string) string {
	wordLength := len(key)
	current := t.root
	stack := []*TrieNode{}
	for i := 0; i < wordLength; i++ {
		index := key[i] - 'a'
		if current.childrens[index] == nil {
			break
		}
		stack = append(stack, current)
		current = current.childrens[index]
	}
	if current.val != "" {
		current.val = ""
	}

	n := len(stack) - 1

	for n > 0 {
		peek := stack[n]
		if peek.val != "" {
			peek = nil
		} else {
			break
		}
		stack = stack[:n]
	}

	return ""
}

func (t *Trie) keys() {

	queue := list.New()
	current := t.root
	collect(current, "", queue)

	for list := queue.Front(); list != nil; list = list.Next() {
		fmt.Printf("%v -> ", list.Value)
	}

}

func (t *Trie) keysWithPrefix(str string) {

	queue := list.New()
	tst := t.GetNode(str)
	collect(tst, str, queue)

	for list := queue.Front(); list != nil; list = list.Next() {
		fmt.Printf("%v -> ", list.Value)
	}

}

func collect(x *TrieNode, prefix string, queue *list.List) {

	if x == nil {
		return
	}
	if x.val != "" {
		queue.PushBack(prefix)
	}

	for c := 0; c < ALBHABET_SIZE; c++ {
		if x.childrens[c] != nil {
			char := string(c + 'a')
			collect(x.childrens[c], prefix+char, queue)
		}
	}

}

func (t *Trie) longestPrefix(str string) {

	length := 0
	wordLength := len(str)
	current := t.root
	if current == nil {
		fmt.Println(str)
		return
	}
	for i := 0; i < wordLength; i++ {
		index := str[i] - 'a'
		if current.childrens[index] != nil {
			if current.val != "" {
				length = i
			}
			current = current.childrens[index]
		}
	}

	fmt.Println(str[:length])

}
