package nombreparfait

import (
	"testing"
)

func Test1(t *testing.T) {
	if !estParfait(6) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if !estParfait(28) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if !estParfait(496) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if !estParfait(8128) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if !estParfait(33550336) {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	for n := uint64(0); n < 10000; n++ {
		if estParfait(n) && n != 6 && n != 28 && n != 496 && n != 8128 {
			t.Fail()
		}
	}
}
