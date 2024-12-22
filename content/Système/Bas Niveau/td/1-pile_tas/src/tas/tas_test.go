package main

import (
	"math/rand"
	"testing"
)

func BenchmarkAdditionBasique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, v := rand.Intn(1000), rand.Intn(1000)
		additionBasique(u, v)
	}
}

func BenchmarkAdditionVersPointeur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, v := rand.Intn(1000), rand.Intn(1000)
		additionVersPointeur(u, v)
	}
}

func BenchmarkAdditionDepuisPointeurs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u, v := rand.Intn(1000), rand.Intn(1000)
		additionDepuisPointeurs(&u, &v)
	}
}
