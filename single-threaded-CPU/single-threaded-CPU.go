package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {

	task := [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}
	fmt.Println(getOrder(task))
}

func getOrder(tasks [][]int) []int {

	cpuTasks := make([]Task, len(tasks))

	for i, task := range tasks {
		var node Task
		node.Index = i
		node.Enqueue = task[0]
		node.Processing = task[1]
		cpuTasks[i] = node
	}

	sort.Slice(cpuTasks, func(a, b int) bool {
		return cpuTasks[a].Enqueue < cpuTasks[b].Enqueue
	})
	fmt.Println("cpuTasks", cpuTasks)
	time := 0
	i := 0
	numProcessed := 0
	result := make([]int, len(cpuTasks))
	tasksPQueue := CpuTaskPQueue{}
	for i < len(cpuTasks) || tasksPQueue.Len() > 0 {

		if tasksPQueue.Len() == 0 && i < len(cpuTasks) && cpuTasks[i].Enqueue > time {
			time = cpuTasks[i].Enqueue
		}
		fmt.Println("time top", time)
		fmt.Println("cpuTasks[i].Enqueue", cpuTasks[i].Enqueue)
		for i < len(cpuTasks) && time >= cpuTasks[i].Enqueue {
			fmt.Println("i", i)
			fmt.Println("cpuTasks[i].Enqueue", cpuTasks[i].Enqueue)
			heap.Push(&tasksPQueue, cpuTasks[i])
			i++
		}

		if tasksPQueue.Len() > 0 {
			fmt.Println("entering if condition")
			task := heap.Pop(&tasksPQueue).(Task)
			time = time + task.Processing
			fmt.Println("task", task)
			fmt.Println("time", time)
			result[numProcessed] = task.Index
			numProcessed++
		}
		fmt.Println("---")
		fmt.Println()
	}

	return result

}

type CpuTaskPQueue []Task

func (c *CpuTaskPQueue) Push(task interface{}) {
	*c = append(*c, task.(Task))
}
func (c *CpuTaskPQueue) Pop() interface{} {
	task := (*c)[len(*c)-1]
	*c = (*c)[:len(*c)-1]
	return task
}
func (c CpuTaskPQueue) Peek() Task {
	return c[0]
}

func (c CpuTaskPQueue) Len() int {
	return len(c)
}

func (c CpuTaskPQueue) Less(a, b int) bool {
	taskA := c[a]
	taskB := c[b]

	if taskA.Processing == taskB.Processing {
		return taskA.Index < taskB.Index
	}
	return taskA.Processing < taskB.Processing
}
func (c CpuTaskPQueue) Swap(a, b int) {
	c[a], c[b] = c[b], c[a]
}

type Task struct {
	Enqueue    int
	Processing int
	Index      int
}
