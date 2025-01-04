# Sparkline library for Go

This is a fork of [https://github.com/dtolpin/gosparkline](David Tolpin's Sparkline Library for Go).

## Example

```Go
import (
	"github.com/janiosarmento/gosparkline"
	"fmt"
)

fmt.Println(spark.Line([]float64{1, 2, 3}))
```

## Changes

Although the entire logic of this library belongs to David, this version has two main differences from the original code:

1. The library now uses a fixed baseline of 0.01 (instead of adjusting the baseline to the minimum value in the series);
2. The code has been split into smaller, specialized functions for clarity and maintainability.

## License

[MIT license](LICENSE).
