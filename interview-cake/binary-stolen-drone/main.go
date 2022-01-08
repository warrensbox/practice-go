package main

import "fmt"

/*

So using booleans probably doesn't save us much space here. Any other ideas?

Let’s zoom out and think about what we’re working with. The only data we have is integers. How are integers stored?

Our machine stores integers as binary numbers ↴ using bits. ↴ What if we thought about this problem on the level of individual bits?

Let's think about the bitwise operations AND, ↴ OR, ↴ XOR, ↴ NOT ↴ and bit shifts. ↴

Is one of those operations helpful for finding our unique integer?

We’re seeing every integer twice, except one. Is there a bitwise operation that would let the second occurrence of an integer cancel out the first?

If so, we could start with a variable uniqueDeliveryId set to 00 and run some bitwise operation with that variable and each number in our array. If duplicate integers cancel each other out, then we’d only be left with the unique integer at the end!

Which bitwise operation would let us do that?

XOR

100
011
111

111
010
101


*/

func main() {

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println("test input")

	uniqueDeliveryId := 0
	deliveryIds := []int{1, 1, 3, 4, 2, 3, 2, 4}

	for _, val := range deliveryIds {
		fmt.Println("val", val)
		uniqueDeliveryId = uniqueDeliveryId ^ val
		fmt.Println(uniqueDeliveryId)
	}

	fmt.Println()
	fmt.Println(uniqueDeliveryId)

}
