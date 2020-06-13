package main

import "fmt"

func main() {
	obj := Constructor(3)
	obj.Set(4, 5)

	param_2 := obj.Snap()
	fmt.Println(param_2)
	obj.Set(3, 6)
	param_3 := obj.Get(3, 0)
	fmt.Println(param_3)
}

type MapArray map[int]int //key: index, val: value

type SnapshotArray struct {
	snapshots map[int]MapArray
	current   MapArray
}

func Constructor(length int) SnapshotArray {
	snapshots := make(map[int]MapArray)
	current := make(MapArray)
	snapshotArr := SnapshotArray{snapshots, current}
	return snapshotArr
}

func (this *SnapshotArray) Set(index int, val int) {
	this.current[index] = val
	fmt.Println("current", this.current)
	return
}

func (this *SnapshotArray) Snap() int {
	snapID := len(this.snapshots)
	fmt.Println("snapID", snapID)
	this.snapshots[snapID] = make(MapArray)
	fmt.Println("snapshots", this.snapshots)
	for k, v := range this.current {
		fmt.Println("k", k)
		fmt.Println("v", v)
		this.snapshots[snapID][k] = v
	}
	fmt.Println("snapshots", this.snapshots)
	return snapID
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return this.snapshots[snap_id][index]
}
