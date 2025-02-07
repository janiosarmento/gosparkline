package spark

// Themes provides predefined color schemes for BarChart.
var Themes = map[string]BarChartConfig{
	"Default": {
		FilledColor: "\033[48;5;30m", // Dark Cyan (Pastel)
		EmptyColor:  "\033[48;5;23m", // Even Darker Cyan
		TextColor:   "\033[97m",      // White text
		ResetColor:  "\033[0m",
	},
	"Classic": {
		FilledColor: "\033[41m", // Red
		EmptyColor:  "\033[44m", // Blue
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Pastel": {
		FilledColor: "\033[48;5;108m", // Soft Green
		EmptyColor:  "\033[48;5;146m", // Soft Blue
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Cyberpunk": {
		FilledColor: "\033[48;5;93m", // Neon Purple
		EmptyColor:  "\033[48;5;21m", // Neon Blue
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Monochrome": {
		FilledColor: "\033[48;5;15m", // White
		EmptyColor:  "\033[48;5;240m", // Gray
		TextColor:   "\033[30m", // Black text
		ResetColor:  "\033[0m",
	},	
	"Solarized": {
		FilledColor: "\033[48;5;136m", // Solarized Yellow
		EmptyColor:  "\033[48;5;33m",  // Solarized Blue
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Dracula": {
		FilledColor: "\033[48;5;92m", // Dracula Green
		EmptyColor:  "\033[48;5;54m", // Dracula Purple
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Gruvbox": {
		FilledColor: "\033[48;5;137m", // Gruvbox Orange
		EmptyColor:  "\033[48;5;94m",  // Gruvbox Brown
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Nord": {
		FilledColor: "\033[48;5;67m", // Nord Blue
		EmptyColor:  "\033[48;5;60m", // Nord Darker Blue
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"OneDark": {
		FilledColor: "\033[48;5;31m", // One Dark Blue
		EmptyColor:  "\033[48;5;236m", // One Dark Background
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"Evergreen": {
		FilledColor: "\033[48;5;29m", // Muted Green
		EmptyColor:  "\033[48;5;22m", // Darker Green
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
	"AutumnGlow": {
		FilledColor: "\033[48;5;137m", // Soft Amber
		EmptyColor:  "\033[48;5;94m",  // Burnt Amber
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	},
}
