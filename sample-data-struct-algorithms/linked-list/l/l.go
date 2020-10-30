package main

//Add : add item to list
func (l *List) Add(item interface{}) {

	var node Node
	node.item = item
	list := l.head

	if list == nil {
		list = &node
	} else {
		for list != nil {
			list = list.next
		}
		list = &node
	}

}

func main() {

	list := List{}

	list.Add(2)

}
