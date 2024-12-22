package souschaine

import "testing"

func Test1(t *testing.T) {
	if sousChaine("", "") != "" {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if sousChaine("bonjour", "") != "" {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if sousChaine("", "au revoir") != "" {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if sousChaine("bonjour", "bonjour") != "bonjour" {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if sousChaine("abcdefghijklmnopqrstuvwxyz", "aeiouyaeiouy") != "aeiouy" {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	if sousChaine("aeiouyaeiouy", "abcdefghijklmnopqrstuvwxyz") != "aeiouy" {
		t.Fail()
	}
}
