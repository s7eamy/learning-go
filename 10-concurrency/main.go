package main

import (
	"fmt"
	"time"
)

func greet(prompt string, num int, done chan map[int]bool) {
	fmt.Println(prompt)
	(<-done)[num] = true
}

func slowGreet(prompt string, num int, done chan map[int]bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(prompt)
	(<-done)[num] = true
	close(done)
}

func main() {
	done := make(chan map[int]bool)
	go greet("Hello from function 1!", 1, done)
	go greet("Hello from function 2!", 2, done)
	go slowGreet("Hello from function 3!", 3, done)
	go greet("Hello from function 4!", 4, done)
	for !((<-done)[1] && (<-done)[2] && (<-done)[3] && (<-done)[4]) {

	}
}
