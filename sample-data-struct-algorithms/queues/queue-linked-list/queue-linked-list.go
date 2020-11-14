package main

func main() {
}

type Node struct {
	item interface{}
	next *Node
}

type Queue struct {
	n    int
	head *Node
	tail *Node
}
