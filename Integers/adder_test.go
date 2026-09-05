package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(69, 67)
	want := 136
	if want != sum {
		t.Errorf("expected %d and got %d", want, sum)
	}
}

func ExampleAdd() {
	a := Add(3, 5)
	fmt.Println(a)
	// Output: 8

}
