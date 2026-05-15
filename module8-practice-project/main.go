package main

import (
	"fmt"
	"module8/practice/project/tax"
)

func main() {

	prices := []float64{10.0, 20.0, 30.0}
	taxRates := []float64{0, 10, 20}

	for _, taxRate := range taxRates {
		job := tax.New(taxRate, prices)
		job.Process()
		fmt.Printf("Tax Rate: %.2f%%\n", taxRate)
		for price, taxIncludedPrice := range job.TaxIncludedPrices {
			fmt.Printf("Original Price: %s, Tax Included Price: %.2f\n", price, taxIncludedPrice)
		}
		fmt.Println()
	}

}
