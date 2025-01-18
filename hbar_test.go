package spark

import "testing"

func TestBarChart(t *testing.T) {
	tests := []struct {
		total   float64
		partial float64
		width   int
		label   string
	}{
		{100, 30, 20, "valor de 30%"},
		{50, 25, 10, "metade"},
		{200, 50, 40, "25% completo"},
		{0, 0, 10, "invalid"}, // Caso total = 0 (valor inválido)
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			result := BarChart(test.total, test.partial, test.width, test.label)
			if len(result) == 0 {
				t.Errorf("Expected a bar chart for %q, got empty string", test.label)
			}
			// Visual validation or pattern matching could go here.
		})
	}
}
