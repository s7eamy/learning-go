package main

import "fmt"

type taxMap map[taxRate][]float64
type taxRate float64

func (t taxRate) applyTax(price float64) float64 {
	return price * (1 + float64(t)/100)
}

func main() {
	var prices []float64 = []float64{10, 20, 30}
	var taxRates []taxRate = []taxRate{0, 7, 15, 25}

	result := make(taxMap, 4)

	for _, taxRate := range taxRates {
		taxedPrices := make([]float64, len(prices))
		for index, price := range prices {
			taxedPrices[index] = taxRate.applyTax(price)
		}
		result[taxRate] = taxedPrices
	}

	fmt.Println("Prices:", prices)
	fmt.Println("Tax rates (%):", taxRates)
	fmt.Println("Taxed prices:", result)
}
