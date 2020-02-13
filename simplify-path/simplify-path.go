package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "/../"
	output := simplifyPath(input)
	fmt.Println(output)
}

func simplifyPath(path string) string {

	var stack []string

	directory := strings.Split(path, "/")

	for _, dir := range directory {

		fmt.Println(dir)
		if dir == ".." && len(stack) > 0 {
			fmt.Println("pop")
			pop(stack)
		} else {
			fmt.Println("push")
			stack = push(stack, dir)
		}
	}

	fmt.Println(stack)

	str := strings.Join(stack[:], "/")

	fmt.Println(str)

	return path

}

// public String simplifyPath(String path) {
//     Deque<String> stack = new LinkedList<>();
//     Set<String> skip = new HashSet<>(Arrays.asList("..",".",""));
//     for (String dir : path.split("/")) {
//         if (dir.equals("..") && !stack.isEmpty()) stack.pop();
//         else if (!skip.contains(dir)) stack.push(dir);
//     }
//     String res = "";
//     for (String dir : stack) res = "/" + dir + res;
//     return res.isEmpty() ? "/" : res;
// }

func push(stack []string, charAt string) []string {

	stack = append(stack, charAt) // Push

	return stack
}

func pop(stack []string) (string, []string) {
	val := ""
	n := len(stack) - 1 // Top element
	val = stack[n]
	stack = stack[:n] // Pop

	fmt.Println("popping")
	fmt.Println(len(stack))
	return val, stack
}
