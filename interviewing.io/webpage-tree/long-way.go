func constructTree(dir []string) {

	graph := make(map[Node][]Node)

	for i := 0; i < len(dir); i++ {
		arr := strings.Split(dir[i], "/")
		graph[Node{"", 0}] = append(graph[Node{"", 0}], Node{arr[0], 1})
		if len(arr) > 0 {
			for j := 1; j < len(arr); j++ {
				parent := arr[j-1]
				child := arr[j]
				graph[Node{parent, j}] = append(graph[Node{parent, j}], Node{child, j + 1})
			}
		}

	}

	var stack []Node
	stack = append(stack, Node{"", 0})
	seen := make(map[Node]bool)

	for len(stack) > 0 {
		size := len(stack)
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if seen[top] {
			continue
		}

		seen[top] = true

		Printer(top)

		for i := 0; i < size; i++ {

			for _, val := range graph[top] {
				if !seen[val] {
					stack = append(stack, val)
				}
			}
		}

	}
}

func Printer(node Node) {

	if node.value == "" {
		return
	}
	indent := strings.Repeat(" ", node.indent)
	value := indent + "--" + node.value
	fmt.Println(value)
}

type Node struct {
	value  string
	indent int
}
