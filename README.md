# Sparkline Library for Go

This is a fork of [Sparkline Library for Go by David Tolpin](https://github.com/dtolpin/gosparkline).

## Example

```Go
import (
	"github.com/janiosarmento/gosparkline"
	"fmt"
)

func main() {
	fmt.Println(spark.Line([]float64{1, 2, 3}))

	// Example using BarChart with default colors
	fmt.Println(spark.BarChart(100, 40, 30, "Memory"))

	// Example using BarChart with custom colors
	config := spark.BarChartConfig{
		FilledColor: "\033[48;5;22m", // Dark Green
		EmptyColor:  "\033[48;5;28m", // Light Green
		TextColor:   "\033[97m",
		ResetColor:  "\033[0m",
	}
	fmt.Println(spark.BarChart(100, 70, 30, "RAM", config))
}
```

## Themes

The library now supports predefined themes for `BarChart`. Instead of manually specifying colors, you can use a theme by passing its name:

```Go
fmt.Println(spark.BarChart(100, 70, 30, "RAM", "Evergreen"))
```

### Available Themes

```Go
Themes["Default"]      // Cyan (default theme)
Themes["Classic"]      // Red and Blue
Themes["Pastel"]       // Soft Green and Blue
Themes["Cyberpunk"]    // Neon Purple and Blue
Themes["Monochrome"]   // Dark Gray and Light Gray
Themes["Solarized"]    // Solarized Yellow and Blue
Themes["Dracula"]      // Green and Purple (Dracula theme)
Themes["Gruvbox"]      // Orange and Brown (Gruvbox theme)
Themes["Nord"]         // Cool Blue and Dark Blue (Nord theme)
Themes["OneDark"]      // One Dark Blue
Themes["Evergreen"]    // Muted Green (RAM theme)
Themes["AutumnGlow"]   // Soft Amber (Swap theme)
```

### Theme Preview  

To see how each theme looks, open [this HTML preview](https://htmlpreview.github.io/?gist.githubusercontent.com/janiosarmento/008d9dabf0c1652639e17e102ce60765/raw/c99b218eebd23ac4048d43ea39c64f5e43b609ee/theme-preview.html) in your browser.  

The preview simulates a dark-themed terminal with all available themes.


## Changes

Although the entire logic of this library belongs to David, this version has several improvements:

1. The library now uses a fixed baseline of 0.01 (instead of adjusting the baseline to the minimum value in the series);
2. The code has been split into smaller, specialized functions for clarity and maintainability;
3. Added `BarChart`, a customizable progress bar with ANSI colors;
4. Users can now define custom color schemes for `BarChart`, or use the default pastel colors;
5. Improved precision in `BarChart` calculations using `math.Round()`;
6. Introduced **predefined themes** for easy styling;
7. Improved unit tests for `BarChart` and `Line()` functions, increasing coverage and robustness.

## Running Tests

To ensure the library is working as expected, you can run the test suite:

```sh
go test ./...
```

## License

[MIT license](LICENSE).
