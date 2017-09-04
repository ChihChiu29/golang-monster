package unittest

import "testing"

func TestAdd(t *testing.T) {
	if Add(3, 4) != 7 {
		t.Error("no!!")
	}
}
