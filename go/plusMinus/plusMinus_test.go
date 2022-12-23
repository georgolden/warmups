package plusMinus

import (
	"testing"
)

func TestPlusMinus(t *testing.T) {
	got := PlusMinus([]int{1, 1, 0, -1, -1})
	want := []float64{0.4, 0.4, 0.2}
	for i, g := range got {
		w := want[i]
		if w != g {
			t.Errorf("Expected '%f', but got '%f'", w, g)
		}
	}
}
