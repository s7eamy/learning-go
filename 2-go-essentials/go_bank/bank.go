package main

import "fmt"

func main() {
	accountBalance := 1000.0

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Print("Choose what you want to do: ")
	var choice uint8
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your account balance:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Money deposit amount: ")
		var amount float64
		fmt.Scan(&amount)
		accountBalance += float64(amount)
		fmt.Println("Success! Your new account balance:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Money withdrawal amount: ")
		var amount float64
		fmt.Scan(&amount)
		accountBalance -= float64(amount)
		fmt.Println("Success! Your new account balance:", accountBalance)
	} else {
		fmt.Println("Thank you for using our services!")
	}
}
