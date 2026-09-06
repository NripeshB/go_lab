package perimeter

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{width: 12, height: 6}, want: 72.0},
		{shape: Circle{radius: 10}, want: 314.1592653589793},
		{shape: Triangle{width: 12, height: 6}, want: 36.0}}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		want := tt.want
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
func Example_floatingComparison() {
	fmt.Println(0.1+0.2 == 0.3)

	var a, b, c float64 = 0.1, 0.2, 0.3
	fmt.Println(a+b == c)

	// Output:
	// true
	// false
}
