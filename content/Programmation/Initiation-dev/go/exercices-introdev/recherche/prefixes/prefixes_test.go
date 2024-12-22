package prefixes

import (
	"testing"
)

func TestTabVide1(t *testing.T) {
	if numPrefixes(nil, "bon") != 0 {
		t.Fail()
	}
}

func TestTabide2(t *testing.T) {
	if numPrefixes([]string{}, "bon") != 0 {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	if numPrefixes([]string{"bonjour", "bonsoir", "salut", "bye bye"}, "bon") != 2 {
		t.Fail()
	}
}

func TestChaineVide1(t *testing.T) {
	if numPrefixes([]string{"bonjour", "bonsoir", "", ""}, "bon") != 2 {
		t.Fail()
	}
}

func TestChaineVide2(t *testing.T) {
	if numPrefixes([]string{"bonjour", "bonsoir", "salut", "bye bye"}, "") != 4 {
		t.Fail()
	}
}

func TestChaineVide3(t *testing.T) {
	if numPrefixes([]string{"bonjour", "bonsoir", "", ""}, "") != 4 {
		t.Fail()
	}
}
