package main

import "sort"

// Ordonnanceur selon la durée de calcul restante (shortest remaining time)
func srt(newTasks []task) (currentTask *task) {
	for i := 0; i < len(newTasks); i++ {
		queue = append(queue, &newTasks[i]) // fill the queue
	}

	sort.Slice(queue, func(i, j int) bool {
		return queue[i].Duration < queue[j].Duration // sort by time remaning
	})

	if queue[0].Duration <= 0 && len(queue) > 0 {
		queue = queue[1:] // if the task is finished delete it
	}

	if len(queue) == 0 {
		return nil // if the queue is empty, the program is finished
	}

	return queue[0]
}
