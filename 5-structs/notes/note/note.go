package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() {
	text := fmt.Sprintf("Note \"%v\" contains content: %v", n.Title, n.Content)
	fmt.Println(text)
}

func (n Note) SaveToFile() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)

	jsonContent, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0444)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("note title and content are both required")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
