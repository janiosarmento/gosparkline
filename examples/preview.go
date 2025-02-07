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
		used := total * (0.4 + r.Float64()*0.4)

		bar := spark.BarChart(total, used, width, themeName, themeName)
		fmt.Println(bar)
		fmt.Println()
	}
}
