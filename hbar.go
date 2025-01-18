package spark

import (
	"strings"
	"unicode/utf8"
)

// BarChart generates a static progress bar with a label overlaid.
// The bar is divided into two sections based on `partial` and `total` values.
// Background colors are red for the filled section and blue for the remaining section.
// Foreground (text) color is always white.
func BarChart(total, partial float64, width int, label string) string {
	if total <= 0 || width <= 0 {
		return ""
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
		bar[i] = ' ' // Red background for the filled section
	}
	for i := partialWidth; i < width; i++ {
		bar[i] = ' ' // Blue background for the unfilled section
	}

	// Overlay the label onto the bar
	for i, r := range labelRunes {
		bar[i] = r
	}

	// Color codes (ANSI escape sequences)
	redBackground := "\033[41;97m" // White text on red background
	blueBackground := "\033[44;97m" // White text on blue background
	reset := "\033[0m"

	// Build the final output with color
	var result strings.Builder
	result.WriteString(redBackground)
	result.WriteString(string(bar[:partialWidth]))
	result.WriteString(reset)
	result.WriteString(blueBackground)
	result.WriteString(string(bar[partialWidth:]))
	result.WriteString(reset)

	return result.String()
}
