package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T){
	t.Run("test1",func(t *testing.T){
		repeated := Repeat("a")
		expected := "aaaaa"
		if repeated != expected{
			fmt.Printf("repeat '%q' but got '%q'",repeated,expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}