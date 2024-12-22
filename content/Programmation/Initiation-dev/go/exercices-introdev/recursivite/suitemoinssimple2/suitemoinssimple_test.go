package suitemoinssimple

import "testing"

func Test5(t *testing.T) {
	res := terme(5)
	if res != 44 {
		t.Error("u(5) vaut 44 mais vous retournez", res)
	}
}

func Test12(t *testing.T) {
	res := terme(12)
	if res != 9 {
		t.Error("u(12) vaut 9 mais vous retournez", res)
	}
}
