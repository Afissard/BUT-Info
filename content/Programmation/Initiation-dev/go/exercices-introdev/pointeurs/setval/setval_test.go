package setval

import "testing"

func Test1(t *testing.T) {
	var a int
	set(&a, 5)
	if a != 5 {
		t.Fail()
	}
}
