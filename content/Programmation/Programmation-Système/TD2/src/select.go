package main

import (
	"fmt"
	"math/rand"
	"time"
)

func routine(c chan int) {
	wait := rand.Intn(1000) // défini un temps d'attente aléatoire
	time.Sleep(time.Duration(wait) * time.Millisecond)
	c <- wait // une fois le temps passé, renvois le temps d’exécution
}

func main() {

	rand.Seed(int64(time.Now().Nanosecond())) // seed

	// créer un channel pour chaque goroutine
	c1 := make(chan int)
	c2 := make(chan int)

	// lance les 2 goroutine
	go routine(c1)
	go routine(c2)

	// affiche la goroutine la plus rapide
	select {
	case w := <-c1:
		fmt.Println("La première goroutine a été la plus rapide.")
		fmt.Println("Elle a mis", w, "millisecondes.")
	case w := <-c2:
		fmt.Println("La deuxième goroutine a été la plus rapide.")
		fmt.Println("Elle a mis", w, "millisecondes.")
	}

}
