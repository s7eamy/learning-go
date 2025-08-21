package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate (%): ")

	earningsBeforeTax, profit, ratio := calculateStats(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInput(prompt string) (input float64) {
	fmt.Print(prompt)
	fmt.Scan(&input)
	return
}

func calculateStats(revenue, expenses, tax float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax/100)
	ratio = ebt / profit
	return
}
