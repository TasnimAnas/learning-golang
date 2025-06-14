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

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid data")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("\n\nNote title: %s\nContent:%s\n\n", note.Title, note.Content)
}

func (note Note) SaveJSON() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = "db/" + strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
