package main

import "fmt"

type Task struct {
	startTime   int
	endTime     int
	parellelism int
}

func main() {
	var taskNum int
	fmt.Scan(&taskNum)
	tasks := []Task{}
	for i := 0; i < taskNum; i++ {
		task := Task{}
		fmt.Scan(&task.startTime, &task.endTime, &task.parellelism)
		tasks = append(tasks, task)
	}
	hashmap := make(map[int]int)
	for _, task := range tasks {
		for i := task.startTime; i < task.endTime; i++ {
			if _, ok := hashmap[i]; !ok {
				hashmap[i] = task.parellelism
			} else {
				hashmap[i] += task.parellelism
			}
		}
	}
	maxPara := 0
	for _, para := range hashmap {
		if para > maxPara {
			maxPara = para
		}
	}
	fmt.Println(maxPara)
}
