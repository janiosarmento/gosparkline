package spark

import (
	"strings"
	"unicode/utf8"
)

// BarChartConfig allows customization of the bar chart colors
type BarChartConfig struct {
	FilledColor string // ANSI escape code for filled section
	EmptyColor  string // ANSI escape code for empty section
	TextColor   string // ANSI escape code for text
	ResetColor  string // ANSI reset code
}

// DefaultConfig returns a default color configuration
func DefaultConfig() BarChartConfig {
	return BarChartConfig{
		FilledColor: "\033[48;5;30m", // Dark Cyan (Pastel)
		EmptyColor:  "\033[48;5;23m", // Even Darker Cyan
		TextColor:   "\033[97m",      // White text
		ResetColor:  "\033[0m",       // Reset
	}
}

// BarChart generates a static progress bar with a customizable color scheme.
// If no custom config is provided, it falls back to the default pastel theme.
func BarChart(total, partial float64, width int, label string, configs ...BarChartConfig) string {
	if total <= 0 || width <= 0 {
		return ""
	}

	// Define the configuration: use user-defined config if provided, otherwise use default
	config := DefaultConfig()
	if len(configs) > 0 {
		config = configs[0]
	}

	// Calculate the widths for the filled and unfilled parts
	partialWidth := int((partial / total) * float64(width))
	if partialWidth > width {
		partialWidth = width
	}

	// Ensure label fits within the bar's width
	labelRunes := []rune(label)
	labelLength := utf8.RuneCountInString(label)
	if labelLength > width {
		labelRunes = labelRunes[:width]
		labelLength = width
	}

	// Create the bar
	bar := make([]rune, width)
	for i := 0; i < partialWidth; i++ {
		bar[i] = ' ' // Filled section
	}
	for i := partialWidth; i < width; i++ {
		bar[i] = ' ' // Empty section
	}

	// Overlay the label onto the bar
	copy(bar, labelRunes)

	// Build the final output with color
	var result strings.Builder
	result.WriteString(config.TextColor)
	result.WriteString(config.FilledColor)
	result.WriteString(string(bar[:partialWidth]))
	result.WriteString(config.ResetColor)
	result.WriteString(config.TextColor)
	result.WriteString(config.EmptyColor)
	result.WriteString(string(bar[partialWidth:]))
	result.WriteString(config.ResetColor)

	return result.String()
}
