package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 12)
	expected := "aaaaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %s but got %s", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 4)
	fmt.Println(repeated)
	// Output: aaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
