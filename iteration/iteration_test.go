package iteration

import (
	"fmt"
	"testing"
)


func TestRepeat(t *testing.T) {

	t.Run("repeat five times", func(t * testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})

	t.Run("repeat ten times", func(t * testing.T) {
		repeated := Repeat("c", 10)
		expected := "cccccccccc"
	
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}		
	})
	
}
func BenchmarkRepeat(b *testing.B){
	for i :=0; i<b.N; i++{
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a",5)
	fmt.Println(repeated)
	// Output: aaaaa
}
