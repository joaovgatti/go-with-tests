package iteration

import "testing"

func TestRepeat(t *testing.T){
	
	
	t.Run("printing the default number of times", func(t *testing.T) {
		repeated := Repeat("a",5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", repeated, expected)
		}
	})

	t.Run("printing the especified number of times", func (t *testing.T){
		repeated := Repeat("a",10)
		expected := "aaaaaaaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", repeated, expected)
		}
	})

}

