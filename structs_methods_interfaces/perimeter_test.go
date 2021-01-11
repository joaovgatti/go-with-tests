package structs_methods_interfaces
import "testing"

func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f hasArea %.2f",got, want)
	}
}

func TestArea(t *testing.T) {


	areaTests := []struct{
		shape Shape
		hasArea float64
	}{
		{Rectangle{12,6},72.0},
		{Circle{10},314.1592653589793},
		{Triangle{12,6},36.0},
	}

	for _,tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %g waant %g",got,tt.hasArea)
		}
	}
}

