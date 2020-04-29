package main

//disclaimer - I have no idea how this works for now
import "fmt"

func main() {

	/**
	 * Your HitCounter object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Hit(timestamp);
	 * param_2 := obj.GetHits(timestamp);
	 */

	obj := Constructor()
	// obj.Hit()
	// obj.Hit(2)
	param_2 := obj.GetHits(301)

	fmt.Println(param_2)
}

const size = 300

type HitCounter struct {
	hits  []int
	total int
	start int
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{
		hits: make([]int, size),
	}
}

func (this *HitCounter) clear(slots int, t int) {
	if slots <= 0 {
		return
	}

	if slots > size {
		slots = size
	}

	for i := 0; i < slots; i++ {
		this.total -= this.hits[this.start%size]
		this.hits[this.start%size] = 0
		this.start = this.start - 1
	}

	if slots == size {
		this.start = t - (t % size)
	}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.clear(timestamp-size-this.start, timestamp)
	this.hits[(timestamp-1)%size]++
	this.total++
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	this.clear(timestamp-size-this.start, timestamp)
	return this.total
}
