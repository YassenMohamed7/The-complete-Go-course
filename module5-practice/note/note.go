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
		return Note{},
			errors.New("title and content must not be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"
	data, err := json.Marshal(n)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)

}

func (n Note) Display() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", n.Title, n.Content, n.CreatedAt.Format("2006-01-02 15:04:05"))
}
