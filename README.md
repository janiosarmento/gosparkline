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

Below is a preview of how each theme looks. The results may vary depending on your terminal color scheme.

<style>
  .bar { width: 300px; height: 20px; position: relative; border-radius: 5px; overflow: hidden; font-family: monospace; text-align: center; line-height: 20px; font-weight: bold; color: white; }
  .bar-fill { height: 100%; position: absolute; top: 0; left: 0; }
</style>

<div class="bar" style="background-color: #003b46;"><div class="bar-fill" style="width: 70%; background-color: #007e7e;">Default</div></div>
<div class="bar" style="background-color: #000080;"><div class="bar-fill" style="width: 70%; background-color: #8b0000;">Classic</div></div>
<div class="bar" style="background-color: #6a8a82;"><div class="bar-fill" style="width: 70%; background-color: #a8d5ba;">Pastel</div></div>
<div class="bar" style="background-color: #240046;"><div class="bar-fill" style="width: 70%; background-color: #5a189a;">Cyberpunk</div></div>
<div class="bar" style="background-color: #444444;"><div class="bar-fill" style="width: 70%; background-color: #bbbbbb;">Monochrome</div></div>
<div class="bar" style="background-color: #073642;"><div class="bar-fill" style="width: 70%; background-color: #b58900;">Solarized</div></div>
<div class="bar" style="background-color: #282a36;"><div class="bar-fill" style="width: 70%; background-color: #50fa7b;">Dracula</div></div>
<div class="bar" style="background-color: #3c3836;"><div class="bar-fill" style="width: 70%; background-color: #d65d0e;">Gruvbox</div></div>
<div class="bar" style="background-color: #2e3440;"><div class="bar-fill" style="width: 70%; background-color: #81a1c1;">Nord</div></div>
<div class="bar" style="background-color: #21252b;"><div class="bar-fill" style="width: 70%; background-color: #61afef;">OneDark</div></div>
<div class="bar" style="background-color: #1e5631;"><div class="bar-fill" style="width: 70%; background-color: #a4de02;">Evergreen</div></div>
<div class="bar" style="background-color: #6b4226;"><div class="bar-fill" style="width: 70%; background-color: #ffb347;">AutumnGlow</div></div>

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
