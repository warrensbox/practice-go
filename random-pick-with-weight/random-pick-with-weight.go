package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Solution struct {
	sums []int
}

func Constructor(w []int) Solution {
	sums := []int{0}
	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
		sums = append(sums, sum)
	}
	return Solution{
		sums: sums,
	}
}

func (this *Solution) PickIndex() int {
	rando := rand.Intn(this.sums[len(this.sums)-1])
	left := 0
	right := len(this.sums) - 1
	for left <= right {
		mid := (left + right) / 2
		if this.sums[mid] <= rando && rando < this.sums[mid+1] {
			return mid
		} else if this.sums[mid] > rando {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -2
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
