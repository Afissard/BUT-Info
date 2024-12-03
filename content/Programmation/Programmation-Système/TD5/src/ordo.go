package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

// Type pour définir une tâche a exécuter
type task struct {
	ID       int // numéro de la tâche
	Priority int // priorité de la tâche (plus elle est haute plus la tâche est proritaire)
	Duration int // durée de la tâche (nombre de pas de temps)
	Arrival  int // temps d'arrivée de la tâche (en pas de temps)
}

var queue []*task = make([]*task, 0)

func main() {

	var fileName string
	var scheduler int

	flag.StringVar(&fileName, "json", "tasks.json", "fichier json contenant la liste des tâches")
	flag.IntVar(&scheduler, "sched", 0, "choix de la fonction d'ordonnancement")
	flag.Parse()

	tasks := getTasks(fileName) // récupération des tâches

	totalDuration := prepareTasks(tasks)     // tri des tâches par ordre d'arrivée
	schedule := make([]int, totalDuration+1) // préparation du planing

	// choix de l'odonnanceur
	var sched func([]task) *task
	switch scheduler {
	case 1:
		sched = sjf
	case 2:
		sched = srt
	case 3:
		sched = rr
	case 4:
		sched = static
	default:
		sched = fcfs
	}

	// Ordonnancement : on demande à l'ordonnanceur, à chaque pas de temps, quelle
	// devrait être la tâche à exécuter, puis on enregistre la réponse et on diminue
	// la durée de cette tâche de 1. On continue ainsi jusqu'à avoir fait un pas
	// de temps de plus que la somme des durées des tâches.
	nextTask := 0
	for t := 0; t < totalDuration+1; t++ {
		from := nextTask
		for nextTask < len(tasks) && tasks[nextTask].Arrival == t {
			nextTask++
		}
		to := nextTask

		runningTask := sched(tasks[from:to])
		if runningTask != nil {
			schedule[t] = runningTask.ID
			runningTask.Duration--
		} else {
			schedule[t] = -1
		}
	}

	displaySchedule(schedule, tasks)

}

// Lecture du fichier json contenant les tâches
func getTasks(fileName string) (tasks []task) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Ouverture json :", err)
	}

	jsonContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Lecture json :", err)
	}

	file.Close()

	err = json.Unmarshal(jsonContent, &tasks)
	if err != nil {
		log.Fatal("Parsing json :", err)
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].ID < tasks[j].ID })

	for i := 0; i < len(tasks)-1; i++ {
		if tasks[i].ID == tasks[i+1].ID {
			log.Fatal("Deux tâches ont la même ID :", tasks[i], tasks[i+1])
		}
	}

	return
}

// Tri des tâches par ordre d'arrivée et calcul du temps total nécessaire à l'exécution
func prepareTasks(tasks []task) (duration int) {

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Arrival < tasks[j].Arrival })

	for _, atask := range tasks {
		duration += atask.Duration
	}

	return
}

// Affichage du résultat de l'ordonnancement dans un navigateur
func displaySchedule(schedule []int, tasks []task) {

	log.Print(schedule)

	htmlpage := fmt.Sprint("<html><body>")

	//for id, s := range sortedSongs {
	for _, aTask := range tasks {
		htmlpage += fmt.Sprint("<p>Task ", aTask.ID, "<canvas id=canvas", aTask.ID, " width=", len(schedule)*20, " height=25></canvas></p>")
		htmlpage += fmt.Sprint("<script>")
		htmlpage += fmt.Sprint("var c = document.getElementById(\"canvas", aTask.ID, "\");")
		htmlpage += fmt.Sprint("var ctx = c.getContext(\"2d\");")
		for time, id := range schedule {
			if id != aTask.ID {
				htmlpage += fmt.Sprint("ctx.fillStyle=\"#ffffff\";")
			} else {
				htmlpage += fmt.Sprint("ctx.fillStyle=\"#00ff00\";")
			}
			htmlpage += fmt.Sprint("ctx.fillRect(", time*20, ",0,", time*20+20, ",20);")
			htmlpage += fmt.Sprint("ctx.beginPath();")
			htmlpage += fmt.Sprint("ctx.moveTo(", time*20, ",0);")
			htmlpage += fmt.Sprint("ctx.lineTo(", time*20, ",20);")
			htmlpage += fmt.Sprint("ctx.closePath();")
			htmlpage += fmt.Sprint("ctx.stroke();")
		}
		htmlpage += fmt.Sprint("</script>")
	}
	htmlpage += fmt.Sprint("</body></html>")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader([]byte(htmlpage)))
	})

	http.ListenAndServe(":8080", nil)

}
