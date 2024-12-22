package souschaine

import "testing"

func Test1(t *testing.T) {
	if !sousChaine("", "") {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if sousChaine("", "a") {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if !sousChaine("a", "") {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if !sousChaine("abcde", "ace") {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if !sousChaine("abcdefghijklmnopqrstuvwxyz", "aeiouy") {
		t.Fail()
	}
}
