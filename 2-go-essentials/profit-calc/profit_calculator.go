package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile string = "financials.txt"

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		processError(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		processError(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate (%): ")
	if err != nil {
		processError(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateStats(revenue, expenses, taxRate)

	results := fmt.Sprintf(`
	EBT: %.2f
	Profit: %.2f
	Ratio: %.2f
	`, earningsBeforeTax, profit, ratio)
	writeToFile(results, resultsFile)
}

func getUserInput(prompt string) (input float64, err error) {
	fmt.Print(prompt)
	fmt.Scan(&input)
	if input <= 0 {
		return 0, errors.New("input must be a positive number")
	}
	return input, nil
}

func processError(err error) {
	fmt.Println("=====")
	fmt.Println("An error occured:", err)
	fmt.Println("Exiting...")
	fmt.Println("=====")
}

func writeToFile(text string, fileName string) {
	err := os.WriteFile(resultsFile, []byte(text), 0644)
	if err != nil {
		processError(err)
	}
}

func calculateStats(revenue, expenses, tax float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax/100)
	ratio = ebt / profit
	return
}
