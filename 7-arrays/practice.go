package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

func main() {
	myHobbies := []string{"Volleyball", "Cybersecurity", "Travelling"}
	fmt.Println("1)")
	additionalHobbies := []string{"Video games", "Reading"}
	myHobbies = append(myHobbies, additionalHobbies...)
	fmt.Println(myHobbies)

	// fmt.Println("2)")
	// fmt.Println(myHobbies[0])
	// fmt.Println([]string{myHobbies[1], myHobbies[2]})
	//
	// fmt.Println("3)")
	// fmt.Println(myHobbies[0:][0:2])
	// fmt.Println(myHobbies[0:2])
	//
	// fmt.Println("4)")
	// fmt.Println(myHobbies[1:2])
	//
	// fmt.Println("5)")
	//
	// goals := []string{"To learn Go syntax", "To learn some Go quirks"}
	// fmt.Println(goals)
	//
	// goals[1] = "A different goal!"
	// goals = append(goals, "My third goal <3")
	// fmt.Println(goals)
	//
	//	products := []product{
	//		{
	//			"Milk",
	//			"PR001",
	//			3.99,
	//		},
	//		{
	//			"Croissant",
	//			"PR002",
	//			0.49,
	//		},
	//	}
	//
	// fmt.Println(products)
	//
	//	products = append(products, product{
	//		"Banana",
	//		"PR003",
	//		0.19,
	//	})
	//
	// fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
