package main

import (
	"fmt"
	"strconv"
)

func main() {

	compute("2-6-7*8/2+5")
}

func compute(express string) {

	index := 0 //define an index value for traversing expressions
	// var num1 uint8
	// var num2 uint8
	// var res uint8 //calculation results
	var num1 int
	var num2 int
	var res int //calculation results
	var ch byte
	var oper byte
	//oper := 0
	numStack := Stack{}
	operStack := Stack{}

	for true {
		ch = express[index] //Continuous traversal operato
		//Determine whether or not it is an operator
		if isOper(ch) {
			//Determine whether the current symbol stack has a symbol
			//if !operStack.isEmpty() {
			//Judging priority if not empty
			//if priority(ch) <= priority(operStack.Peek().(byte)) {
			//When the priority is less than the value at the top of the stack, two stack values pop up for calculation.
			// num1 = numStack.Pop().(int)
			// num2 = numStack.Pop().(int)
			// oper = operStack.Pop().(byte)
			// res = cal(num1, num2, oper)

			for numStack.Len() >= 2 && operStack.Len() >= 1 {
				if priority(ch) <= priority(operStack.Peek().(byte)) {
					num1 = numStack.Pop().(int)
					num2 = numStack.Pop().(int)
					oper = operStack.Pop().(byte)
					res = cal(num1, num2, oper)
					numStack.Push(res)
				} else {
					break
				}
			}

			//After calculation, the calculated value is put on the stack of numbers.
			fmt.Println("res", res)
			//numStack.Push(res)
			fmt.Println(numStack.stack...)
			//At the same time, put the operators in the symbol stack.
			fmt.Println(string(ch))
			//if priority(ch) <= priority(operStack.Peek().(byte)) {
			operStack.Push(ch)
			//}
			printStack(operStack.stack)
			// } else {
			// 	//priority
			// 	fmt.Println("priority", string(ch))
			// 	operStack.Push(ch)
			// 	printStack(operStack.stack)
			// }
			// } else {
			// 	fmt.Println("empty", string(ch))
			// 	//To stack symbols directly for emptiness
			// 	operStack.Push(ch)
			// 	printStack(operStack.stack)
			// }
		} else {
			//fmt.Println("ch", string(ch))
			intVar, _ := strconv.Atoi(string(ch))
			numStack.Push(intVar)
			fmt.Println(numStack.stack...)
		}
		index++
		if index >= len(express) {
			break
		}
	}

	//After scanning, the value of the stack is calculated with the value in the operator.
	for true {
		if operStack.isEmpty() {
			break
		}
		num1 = numStack.Pop().(int)
		num2 = numStack.Pop().(int)
		fmt.Println("NUM1", num1)
		fmt.Println("NUM2", num2)
		oper = operStack.Pop().(byte)
		res = cal(num1, num2, oper)
		numStack.Push(res)
		fmt.Println(numStack.stack...)
	}
	fmt.Println(numStack.Pop().(int))
}
func printStack(arr []interface{}) {
	arr2 := []string{}
	for _, v := range arr {
		//fmt.Println(string(v.(uint8)))
		arr2 = append(arr2, string(v.(uint8)))
	}
	fmt.Println(arr2)
}

//Determine whether it is an operator
//@param oper  Incoming characters
//@return  If the operator returns true, otherwise it returns false
func isOper(oper byte) bool {
	return oper == '+' || oper == '-' || oper == '*' || oper == '/'
}

/**
* Judging the priority of operators
* @param oper Priority of incoming
* @return  The return priority is 1, -1, 0, respectively.
 */
func priority(oper byte) int {
	if oper == '*' || oper == '/' {
		return 2
	} else if oper == '+' {
		return 1
	} else if oper == '-' {
		return 0
	} else {
		return -1
	}
}

//computing method
func cal(num1, num2 int, oper byte) int {
	var res int
	switch oper {
	case '+':
		fmt.Println("num1+", num1)
		fmt.Println("num2+", num2)
		res = num1 + num2
		break
	case '-':

		fmt.Printf("%v - %v\n", num2, num1)
		res = num2 - num1
		break
	case '*':
		res = num1 * num2
		break
	case '/':
		fmt.Println("num1/", num1)
		fmt.Println("num2/", num2)
		res = num2 / num1
	}
	return res
}

/***********************************/

type Stack struct {
	stack []interface{}
}

func (s *Stack) isEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(item interface{}) {

	s.stack = append(s.stack, item)

}

func (s *Stack) Pop() interface{} {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]
	s.stack = s.stack[:n] //pop

	return top

}

func (s *Stack) Peek() interface{} {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]

	return top

}

func (s *Stack) Len() int {

	return len(s.stack)

}
