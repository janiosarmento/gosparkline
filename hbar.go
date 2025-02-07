package spark

import (
	"math"
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
	return Themes["Default"]
}

// BarChart generates a static progress bar with a customizable color scheme or a predefined theme.
// If no custom config is provided, it falls back to the default theme.
func BarChart(total, partial float64, width int, label string, themeOrConfig interface{}) string {
	if total <= 0 || width <= 0 {
		return ""
	}

	var config BarChartConfig

	switch v := themeOrConfig.(type) {
	case string:
		if theme, exists := Themes[v]; exists {
			config = theme
		} else {
			config = DefaultConfig()
		}
	case BarChartConfig:
		config = v
	default:
		config = DefaultConfig()
	}

	// Calculate the widths for the filled and unfilled parts
	partialWidth := int(math.Round((partial / total) * float64(width)))
	if partialWidth > width {
		partialWidth = width
	}

	// Ensure label fits within the bar's width
	labelRunes := []rune(label)
	if utf8.RuneCountInString(label) > width {
		labelRunes = labelRunes[:width]
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
