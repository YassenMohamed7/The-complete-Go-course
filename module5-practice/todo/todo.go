package todo

import (
	"encoding/json"
	"os"
)

type Todo struct {
	Text string `json:"todo_text"`
}

func New(text string) Todo {
	return Todo{Text: text}
}

func (t Todo) Display() {
	println(`Todo text : ` + t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
