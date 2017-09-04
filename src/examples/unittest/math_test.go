package unittest

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(3, 4) != 7 {
		t.Error("no!!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(3, 4)
	}
}

func ExampleAdd() {
	fmt.Println(Add(3, 4))
	// Output: 7
}
