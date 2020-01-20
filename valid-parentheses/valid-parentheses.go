package main

import "fmt"


func main(){

	ans := isValid("{[{[]()}]}")
	fmt.Println(ans)
}


func isValid(s string) bool {


	var stack []string

	dict := make(map[string]string)
	dict[")"] = "("
	dict["}"] = "{"
	dict["]"] = "["

    for _,val := range s {
		fmt.Println("type", string(val))
		if _, ok :=  dict[string(val)]; ok {
		
			// Pop the topmost element from the stack, if it is non empty
		    // Otherwise assign a dummy value of '#' to the topElement variable
			//topElement = stack.pop() if stack else '#'
			topElement := ""
			if len(stack) > 0 {
				topElement,stack = pop(stack)
			}else{
				topElement = "#"
			}
			fmt.Println("topElement", topElement)
			fmt.Println("dict[string(val)]", dict[string(val)])
			if dict[string(val)] != topElement {
				fmt.Println("gohere")
				return false
			}
			 fmt.Println(topElement)
		}else{
			stack = push(stack,string(val))
			fmt.Println(stack)
		}
		fmt.Println(len(stack))
		fmt.Println("============")
	}
	


	if len(stack) > 0 {
		return false
	}
	
	return true

}

func push(stack []string, charAt string) []string {

	stack = append(stack,charAt) // Push

	return stack
}

func pop(stack []string) (string, []string){
	val := ""
	n := len(stack) - 1 // Top element
	val = stack[n]
	stack = stack[:n] // Pop

	fmt.Println("popping")
	fmt.Println(len(stack))
	return val, stack
}


