package main

import "sort"

var prevQueueLen = 0
var aTaskHaveFinished = false

// Ordonnanceur selon la durée de calcul (shortest job first)
func sjf(newTasks []task) (currentTask *task) {
	prevQueueLen = len(queue)
	for i := 0; i < len(newTasks); i++ {
		queue = append(queue, &newTasks[i]) // fill the queue
	}

	if prevQueueLen != len(queue) || aTaskHaveFinished { // sort if a new job is added
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Duration < queue[j].Duration
		})
	}

	if queue[0].Duration <= 0 && len(queue) > 0 {
		queue = queue[1:] // if the task is finished delete it
		aTaskHaveFinished = true
	} else {
		aTaskHaveFinished = false
	}

	if len(queue) == 0 {
		return nil // if the queue is empty, the program is finished
	}

	return queue[0]
}
