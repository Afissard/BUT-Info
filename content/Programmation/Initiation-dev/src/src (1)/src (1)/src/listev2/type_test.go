package liste

import "testing"

func BenchmarkCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := Vide()
		for j := 0; j < 100000; j++ {
			l = Concat(j, l)
		}
	}
}
