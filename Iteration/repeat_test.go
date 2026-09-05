package iteration

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 5)
	want := "aaaaa"

	if want != repeated {
		t.Errorf("Expected %q, but got %q", want, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
