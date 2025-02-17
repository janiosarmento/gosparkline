// Sparklines
package spark // import "github.com/janiosarmento/gosparkline"

import (
	"math"
)

var levels = []rune("▁▂▃▄▅▆▇█")

// Line generates a sparkline string from a slice of float64.
func Line(ys []float64) string {
	if len(ys) == 0 {
		return ""
	}

	const base = 0.01
	const upper_min = 1.0
	
	max := FindMax(ys)

	if max <= base { // Garantir escala válida
		max = base + 1
	}

	if max < upper_min {
		max = upper_min
	}

	return generateSparkline(ys, base, max)
}
// generateSparkline maps values to rune levels and creates a sparkline string.
func generateSparkline(ys []float64, base, max float64) string {
	line := make([]rune, len(ys))
	scale := float64(len(levels) - 1)

	for i, y := range ys {
		j := int(math.Floor(scale * (y - base) / (max - base)))
		if j < 0 {
			j = 0
		} else if j >= len(levels) {
			j = len(levels) - 1
		}
		line[i] = levels[j]
	}

	return string(line)
}

// findMax returns the maximum value in a slice of float64.
func FindMax(ys []float64) float64 {
	max := math.Inf(-1)
	for _, y := range ys {
		if y > max {
			max = y
		}
	}
	return max
} 