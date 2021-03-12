package main

import (
	"container/list"
	"fmt"
	"math"
)

/**
 * Problem:
 * Kth multiple number;
 * Kth number with only prime factors of 3, 5 and 7.
 * ;
 * ;
 * Fact:
 * Kth Multiple of 3, 5 and 7 = 3^k * 5^k * 7^k
 * ;
 * ;
 * Solution:
 * - Take 3 queues
 * - - Queue-3 to hold 3 multiples
 * - - Queue-5 to hold 5 multiples
 * - - Queue-7 to hold 7 multiples
 * - Remove minimum value from all three queues
 * - If we removed from Q3 then
 * - - add Q3 <- min * 3, Q5 <- min * 5, Q7 <- min * 7
 * - If we removed from Q5 then
 * - - add Q5 <- min * 5, Q7 <- min * 7
 * - If we removed from Q7 then
 * - - add Q7 <- min * 7
 * - - We'll not add to Q3 and Q5 because it will be duplicate
 * ;
 * ;
 * Algorithm:
 * - Take three queue namely Q3, Q5 and Q7
 * - Add 1 to Q3
 * - Iterate 1 to K
 * - - Fetch minimum_value from all three Queues
 * - - If minimum_value == Q3.peek() then
 * - - - Q3.removeMin()
 * - - - Add Q3 <- minimum_value * 3, Q5 <- minimum_value * 5 and Q7 <- minimum_value * 7
 * - - If minimum_value == Q5.peek() then
 * - - - Q5.removeMin()
 * - - - Add Q5 <- minimum_value * 5 and Q7 <- minimum_value * 7
 * - - If minimum_value == Q7.peek() then
 * - - - Q7.removeMin()
 * - - - Add Q7 <- minimum_value * 7
 * - return minimum_value
 *
 */

func main() {
	fmt.Println(allPossibleKFactors(3))
}

func allPossibleKFactors(k int) int {

	q3 := list.New()
	q5 := list.New()
	q7 := list.New()
	var v3, v5, v7 int
	val := 0
	q3.PushBack(1)

	for i := 0; i <= k; i++ {

		if q3.Len() > 0 {
			v3 = q3.Front().Value.(int) //cast
		} else {
			v3 = math.MaxInt16
		}
		if q5.Len() > 0 {
			v5 = q5.Front().Value.(int) //cast
		} else {
			v5 = math.MaxInt16
		}
		if q7.Len() > 0 {
			v7 = q7.Front().Value.(int) //cast
		} else {
			v7 = math.MaxInt16
		}

		val = Min(v3, Min(v5, v7))

		if val == v3 {
			elem := q3.Front() //dequeue
			q3.Remove(elem)
			q3.PushBack(val * 3)
			q5.PushBack(val * 5)
		} else if val == v5 {
			elem := q5.Front() //dequeue
			q5.Remove(elem)
			q5.PushBack(val * 5)
		} else if val == v7 {
			elem := q7.Front() //dequeue
			q7.Remove(elem)
		}

		q7.PushBack(val * 7)

	}

	return val
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
