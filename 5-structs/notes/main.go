package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	myNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	myNote.Display()

	err = myNote.SaveToFile()
	if err != nil {
		fmt.Println("Failed to save to file:", err)
		return
	}
	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	input = handleLongInput()
	return input
}

func handleLongInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}
