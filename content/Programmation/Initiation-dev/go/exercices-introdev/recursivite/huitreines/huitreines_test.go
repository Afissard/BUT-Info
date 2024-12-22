package huitreines

import "testing"

func TestPasSolution1(t *testing.T) {
	_, ok := huitreines(2)
	if ok {
		t.Fail()
	}
}

func TestPasSolution2(t *testing.T) {
	_, ok := huitreines(3)
	if ok {
		t.Fail()
	}
}

func TestSolution1(t *testing.T) {
	sol, ok := huitreines(1)
	if !ok || len(sol) < 1 || len(sol[0]) < 1 || sol[0][0] != 1 {
		t.Fail()
	}
}

func TestSolution2(t *testing.T) {
	for n := 4; n <= 10; n++ {
		sol, ok := huitreines(n)
		if !ok {
			t.Fail()
		} else if !bonnesDimensions(sol, n) {
			t.Fail()
		} else if numReines(sol) != n {
			t.Fail()
		} else if !estSansConflits(sol) {
			t.Fail()
		}
	}
}

// Fonctions utilitaires pour les tests
func estSansConflits(plateau [][]int) bool {
	for l := 0; l < len(plateau); l++ {
		numReines := 0
		posReine := -1
		for c := 0; c < len(plateau[l]); c++ {
			numReines += plateau[l][c]
			if plateau[l][c] == 1 {
				posReine = c
			}
		}
		if numReines > 1 {
			return false
		}
		if posReine >= 0 {
			for ll := 0; ll < len(plateau); ll++ {
				if ll != l {
					if plateau[ll][posReine] == 1 {
						return false
					}
					for cc := 0; cc < len(plateau[ll]); cc++ {
						if plateau[ll][cc] == 1 {
							diffL := ll - l
							diffC := cc - posReine
							if diffL == diffC || diffL == -diffC {
								return false
							}
						}
					}
				}
			}
		}
	}
	return true
}

func bonnesDimensions(plateau [][]int, n int) bool {
	if len(plateau) != n {
		return false
	}
	for i := 0; i < n; i++ {
		if len(plateau[i]) != n {
			return false
		}
	}
	return true
}

func numReines(plateau [][]int) (num int) {
	for i := 0; i < len(plateau); i++ {
		for j := 0; j < len(plateau[i]); j++ {
			num += plateau[i][j]
		}
	}
	return num
}
