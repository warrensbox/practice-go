package main

import "fmt"

type RLEIterator struct {
	index  int
	ithNum int
	queue  []int
}

func Constructor(encoding []int) RLEIterator {

	var rlei RLEIterator
	rlei.queue = encoding
	return rlei
}

func (this *RLEIterator) Next(n int) int {
	//fmt.Println("this.ithNum",this.ithNum)
	//fmt.Println("-------")
	for n > 0 && this.ithNum < len(this.queue) {
		fmt.Println("this.NUM1", this.ithNum)

		fmt.Println(" this.index=1", this.index)
		fmt.Println(" n=1", n)
		fmt.Println("this.queue[this.ithNum]=1", this.queue[this.ithNum])
		if this.queue[this.ithNum]-this.index >= n {
			//fmt.Println(" this.indexXXX", this.index)
			this.index += n
			// fmt.Println(" this.indexXXX", this.index)
			// fmt.Println(" this.queue[this.ithNum+1]",this.queue[this.ithNum+1])
			fmt.Println("returning")
			fmt.Println()
			return this.queue[this.ithNum+1]
		}
		fmt.Println()
		fmt.Println(" going here n ", n)
		//fmt.Println("this.NUUMM222",this.ithNum)
		fmt.Println("this.index222", this.index)
		fmt.Println("this.queue[this.ithNum]", this.queue[this.ithNum])
		n = n - (this.queue[this.ithNum] - this.index)
		fmt.Println("nuu", n)
		this.index = 0
		this.ithNum += 2

	}

	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */
