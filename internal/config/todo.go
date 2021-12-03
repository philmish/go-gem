package config

import "fmt"

type Todo struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
	Urgency int8   `json:"urgency"`
	Id      int    `json:"id"`
}

func newTodo(content string, urgency int8, id int) *Todo {
	return &Todo{content, false, urgency, id}
}

func (t *Todo) formatTodo() string {
	return fmt.Sprintf("(%d)%s\nID: %d\n\n", t.Urgency, t.Content, t.Id)
}
