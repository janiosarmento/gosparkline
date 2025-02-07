package spark

import (
	"strings"
	"testing"
)

func TestBarChart_Default(t *testing.T) {
	total := 100.0
	partial := 50.0
	width := 20
	label := "Test"

	bar := BarChart(total, partial, width, label, "Default")

	if !strings.Contains(bar, "Test") {
		t.Errorf("Expected label 'Test' in output, got: %s", bar)
	}
	if len(bar) == 0 {
		t.Errorf("Expected non-empty output, got empty string")
	}
}

func TestBarChart_CustomTheme(t *testing.T) {
	total := 100.0
	partial := 75.0
	width := 30
	label := "Custom"

	bar := BarChart(total, partial, width, label, "Evergreen")

	if !strings.Contains(bar, "Custom") {
		t.Errorf("Expected label 'Custom' in output, got: %s", bar)
	}
	if len(bar) == 0 {
		t.Errorf("Expected non-empty output, got empty string")
	}
}

func TestBarChart_InvalidThemeFallback(t *testing.T) {
	total := 100.0
	partial := 30.0
	width := 25
	label := "Fallback"

	bar := BarChart(total, partial, width, label, "NonExistentTheme")

	if !strings.Contains(bar, "Fallback") {
		t.Errorf("Expected label 'Fallback' in output, got: %s", bar)
	}
	if len(bar) == 0 {
		t.Errorf("Expected non-empty output, got empty string")
	}
}
