package spark

import (
	"testing"
	"unicode/utf8"
)

func TestLine(t *testing.T) {
	for _, c := range []struct {
		name   string
		ys     []float64
		length int
	}{
		{"empty", []float64{}, 0},
		{"regular", []float64{1, 3, 2}, 3},
		{"single element", []float64{1}, 1},
		{"constant series", []float64{3, 3, 3}, 3},
		{"increasing series", []float64{1, 2, 3, 4, 5}, 5},
		{"decreasing series", []float64{5, 4, 3, 2, 1}, 5},
		{"contains zero", []float64{0, 1, 2, 3}, 4},
		{"contains negatives", []float64{-1, 0, 1, 2}, 4},
		{"large values", []float64{1000, 2000, 3000, 4000}, 4},
	} {
		line := Line(c.ys)
		length := utf8.RuneCountInString(line)
		if length != c.length {
			t.Errorf("%s: expected sparkline of length %d, got %q of length %d",
				c.name, c.length, line, length)
		}
	}
}
