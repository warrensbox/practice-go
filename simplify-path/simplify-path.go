package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	input := "/home/run"
	output := simplifyPath(input)
	fmt.Println(output)
}

func simplifyPath(path string) string {

	//var stack []string

	dir, file := filepath.Split(path)

	fmt.Println(dir)
	fmt.Println(file)
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
