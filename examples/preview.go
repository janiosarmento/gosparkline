package main

import (
	"fmt"
	"math/rand"
	"time"

	spark "github.com/janiosarmento/gosparkline"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	total := 100.0
	width := 50

	fmt.Println("\n=== BarChart Theme Preview ===")
	fmt.Println()

	for themeName := range spark.Themes {
		used := total * (0.2 + r.Float64()*0.6)

		bar := spark.BarChart(total, used, width, themeName, themeName)
		fmt.Println(bar)
		fmt.Println()
	}

	fmt.Println("\n=== Sparkline Preview ===")
	fmt.Println()

	values := make([]float64, width)
	for i := range values {
		values[i] = r.Float64() * 100
	}
	sparkline := spark.Line(values)
	fmt.Println(sparkline)
	fmt.Println()
}