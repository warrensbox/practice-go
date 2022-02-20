package main

import "container/list"

func main() {

}

type LRUCache struct {
	LinkedList *list.List
	Cache      map[int]*list.Element
	Capacity   int
}

func Constructor(capacity int) LRUCache {
	//init list, cache
	return LRUCache{
		LinkedList: list.New(),
		Cache:      make(map[int]*list.Element),
		Capacity:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	//if it's not in the cache, return -1
	//get the element from the cache
	//move it to the front of the list
	//return element.Value
	elem, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.LinkedList.MoveToFront(elem)
	return elem.Value.([]int)[1] //0 is the key and 1 is the value
}

func (this *LRUCache) Put(key int, value int) {
	//check if element is already in cache, if yes, update cache with new value
	//move element to the front
	elem, ok := this.Cache[key]
	if ok {
		this.LinkedList.Remove(elem) //rm the existing elem- you can just move it to the front, you have to delete it
		new_elem := this.LinkedList.PushFront([]int{key, value})
		this.Cache[key] = new_elem
		return
	}

	//if the capacity has been reached,  get the last element at the back of the list
	//remove it
	if len(this.Cache) == this.Capacity {
		last_elem := this.LinkedList.Back()
		this.LinkedList.Remove(last_elem)
		delete(this.Cache, last_elem.Value.([]int)[0])
	}

	//delete the element form the cache
	//create new element, append to the front of the list, add to cache
	new_elem := this.LinkedList.PushFront([]int{key, value})
	this.Cache[key] = new_elem
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
