package iteration

import "testing"


func TestRepeat(t *testing.T) {

	t.Run("repeat five times", func(t * testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})

	t.Run("repeat six times", func(t * testing.T) {
		repeated := Repeat("a")
		expected := "aaaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})
}