package main

import (
	"fmt"
	"strconv"
)

func main() {
	compute("2-6-7*8/2+5")
}

func compute(express string) {

	var num1 int
	var num2 int
	var res int //calculation results
	var ch byte
	var oper byte
	numStack := Stack{}
	operStack := Stack{}

	for i := 0; i < len(express); i++ {
		ch = express[i] //Continuous traversal operation
		intVar, _ := strconv.Atoi(string(ch))
		numStack.Push(intVar)
		i++
		if i >= len(express) {
			break
		}
		ch = express[i] //Continuous traversal operation

		if isOper(ch) {
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
			operStack.Push(ch)
		}
	}

	for true {
		if operStack.isEmpty() {
			break
		}
		num1 = numStack.Pop().(int)
		num2 = numStack.Pop().(int)
		oper = operStack.Pop().(byte)
		res = cal(num1, num2, oper)
		numStack.Push(res)
		fmt.Println(numStack.stack...)
	}

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
	} else if oper == '-' || oper == '+' {
		return 1
	} else {
		return -1
	}
}

//computing method
func cal(num1, num2 int, oper byte) int {
	var res int
	switch oper {
	case '+':
		res = num1 + num2
	case '-':
		res = num2 - num1
	case '*':
		res = num1 * num2
	case '/':
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
		return -1
	}
	n := len(s.stack) - 1
	top := s.stack[n]
	s.stack = s.stack[:n] //pop
	return top
}

func (s *Stack) Peek() interface{} {
	if len(s.stack) == 0 {
		return -1
	}
	n := len(s.stack) - 1
	top := s.stack[n]
	return top
}

func (s *Stack) Len() int {
	return len(s.stack)
}
