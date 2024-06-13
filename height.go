package main

import (
	"flag"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	var height float64
	flag.Float64Var(&height, "height", 5.6, "Input Height")
	flag.Parse()
	matrixValue := decimal.NewFromFloat(convertFeetToMeters(height)).Round(2)
	fmt.Printf("Height in centimeter %v\n", matrixValue)
}

func convertFeetToMeters(feet float64) float64 {
	tall := 0.3048 * feet
	return tall
}
