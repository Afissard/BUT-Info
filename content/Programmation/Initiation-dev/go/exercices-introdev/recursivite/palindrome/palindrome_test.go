package palindrome

import (
	"testing"
)

func TestVide(t *testing.T) {
	if !estPalindrome("") {
		t.Fail()
	}
}

func TestPalindrome1(t *testing.T) {
	if !estPalindrome("a") {
		t.Fail()
	}
}

func TestPalindrome2(t *testing.T) {
	if !estPalindrome("lol") {
		t.Fail()
	}
}

func TestPalindrome3(t *testing.T) {
	if !estPalindrome("kayak") {
		t.Fail()
	}
}

func TestPalindrome4(t *testing.T) {
	if !estPalindrome("rionsnoir") {
		t.Fail()
	}
}

func TestPalindrome5(t *testing.T) {
	if !estPalindrome("engagelejeuquejelegagne") {
		t.Fail()
	}
}

func TestPasPalindrome1(t *testing.T) {
	if estPalindrome("ab") {
		t.Fail()
	}
}

func TestPasPalindrome2(t *testing.T) {
	if estPalindrome("bonjour") {
		t.Fail()
	}
}

func TestPasPalindrome3(t *testing.T) {
	if estPalindrome("kayaks") {
		t.Fail()
	}
}
