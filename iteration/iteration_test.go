package iteration

import (
	"fmt"
	"testing"
)


func TestRepeat(t *testing.T) {

	t.Run("repeat five times", func(t * testing.T) {
		repeated := Repeat("a", 5, false)
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})

	t.Run("repeat ten times", func(t * testing.T) {
		repeated := Repeat("c", 10, false)
		expected := "cccccccccc"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})

	t.Run("uses repeat from string lib", func(t * testing.T) {
		repeated := Repeat("ba", 3, true)
		expected := "bababa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})
	
}
func BenchmarkRepeat(b *testing.B){
	for i :=0; i<b.N; i++{
		Repeat("a", 5, false)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a",5, false)
	fmt.Println(repeated)
	// Output: aaaaa
}
