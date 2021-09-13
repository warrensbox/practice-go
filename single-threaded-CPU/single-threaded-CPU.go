package main

import (
	"container/heap"
	"sort"
)

func getOrder(tasks [][]int) []int {

	//var res []int
	//var process Process
	//process := make(Process,len(tasks))
	// for i, task := range tasks {
	//     var node Task
	//     node.Index = i
	//     node.Enqueue = task[0]
	//     node.Processing = task[1]
	//     process = append(process,node)
	// }
	cpuTasks := make([]Task, len(tasks))

	for i, task := range tasks {
		var node Task
		node.Index = i
		node.Enqueue = task[0]
		node.Processing = task[1]
		cpuTasks[i] = node
	}

	//fmt.Println(process)
	//sort.Sort(SortTask(cpuTasks))
	//sort.SliceStable(cpuTasks, func(i, j int) bool { return cpuTasks[i].Enqueue < cpuTasks[j].Enqueue })
	sort.Slice(cpuTasks, func(a, b int) bool {
		return cpuTasks[a].Enqueue < cpuTasks[b].Enqueue
	})
	//fmt.Println(cpuTasks)
	// sort.SliceStable(tasks, func(i, j int) bool { return tasks[i][0] < tasks[j][0] })
	//fmt.Println(tasks)

	//var stack Process
	time := 0
	i := 0
	numProcessed := 0
	result := make([]int, len(cpuTasks))
	tasksPQueue := CpuTaskPQueue{}
	for i < len(cpuTasks) || tasksPQueue.Len() > 0 {

		if tasksPQueue.Len() == 0 && i < len(cpuTasks) && cpuTasks[i].Enqueue > time {
			time = cpuTasks[i].Enqueue
			//continue
		}

		for i < len(cpuTasks) && time >= cpuTasks[i].Enqueue {

			//stack = append(stack,cpuTasks[i])
			heap.Push(&tasksPQueue, cpuTasks[i])
			i++
		}
		//fmt.Println("stack",stack)
		//sort.Sort(Process(stack))
		if tasksPQueue.Len() > 0 {
			task := heap.Pop(&tasksPQueue).(Task)
			//top := stack[len(stack)-1]
			//stack = stack[:len(stack)-1]
			time = time + task.Processing
			//res = append(res,top.Index)
			result[numProcessed] = task.Index
			numProcessed++
		}
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

// type Process []Task
// type SortTask []Task

// func (this Process) Less(i, j int) bool {
//     if this[i].Processing == this[j].Processing {
//         return this[i].Index > this[j].Index
//     }
//     return this[i].Processing > this[j].Processing
// }
// func (this Process) Len() int { return len(this)}
// func (this Process) Swap(i, j int) { this[i],this[j] = this[j],this[i] }

// func (this SortTask) Less(i, j int) bool { return this[i].Enqueue < this[j].Enqueue }
// func (this SortTask) Len() int { return len(this)}
// func (this SortTask) Swap(i, j int) { this[i],this[j] = this[j],this[i] }

type Task struct {
	Enqueue    int
	Processing int
	Index      int
}
