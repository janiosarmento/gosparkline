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

## Changes

Although the entire logic of this library belongs to David, this version has several improvements:

1. The library now uses a fixed baseline of 0.01 (instead of adjusting the baseline to the minimum value in the series);
2. The code has been split into smaller, specialized functions for clarity and maintainability;
3. Added `BarChart`, a customizable progress bar with ANSI colors;
4. Users can now define custom color schemes for `BarChart`, or use the default pastel colors;
5. Improved precision in `BarChart` calculations using `math.Round()`.

## Contributing

Contributions are welcome! If you find a bug, have a feature request, or want to improve the library, feel free to open an issue or submit a pull request on [GitHub](https://github.com/janiosarmento/gosparkline).

Before submitting a PR, please:
- Follow the existing code style and structure.
- Write clear commit messages.
- Test your changes thoroughly.

## License

[MIT license](LICENSE).
