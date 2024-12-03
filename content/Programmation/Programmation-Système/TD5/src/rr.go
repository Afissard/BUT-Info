package main

import (
	"sort"
)

const QUANTUM = 5

var (
	quantumCount  = 0
	currentTaskId = 0
)

func getTaskById(id int) *task {
	for i := 0; i < len(queue); i++ {
		if queue[i].ID == id {
			return queue[i]
		}
	}
	// log.Fatal("nop")
	return nil
}

// Ordonnanceur à temps partagé avec politique du tourniquet (Round Robin)
// le quantum vaudra 5 unités de temps
func rr(newTasks []task) (currentTask *task) {
	// init
	for i := 0; i < len(newTasks); i++ {
		queue = append(queue, &newTasks[i]) // fill the queue
	}

	sort.Slice(queue, func(i, j int) bool {
		return queue[i].ID < queue[j].ID // sort by priority
	})

	// algo
	if quantumCount > 0 {
		quantumCount--
	} else {
		quantumCount = QUANTUM
		currentTaskId = queue[currentTaskId+1%len(queue)].ID
		return rr(newTasks)
	}

	// returning value
	if queue[0].Duration <= 0 && len(queue) > 0 {
		queue = queue[1:] // if the task is finished delete it
	}

	if len(queue) == 0 {
		return nil // if the queue is empty, the program is finished
	}

	return getTaskById(currentTaskId)
}
