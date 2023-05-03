package main

import (
	"fmt"
	"time"
)

type TodoItem struct {
	Text string
	ID   int64
}

type TodoApp struct {
	Items []TodoItem
	Text  string
}

func (t *TodoApp) handleChange(text string) {
	t.Text = text
}

func (t *TodoApp) handleSubmit() {
	if len(t.Text) == 0 {
		return
	}
	newItem := TodoItem{
		Text: t.Text,
		ID:   time.Now().Unix(),
	}
	t.Items = append(t.Items, newItem)
	t.Text = ""
}

type TodoList struct {
	Items []TodoItem
}

func main() {
	items := []TodoItem{}
	text := ""

	todoApp := TodoApp{
		Items: items,
		Text:  text,
	}

	todoList := TodoList{
		Items: items,
	}

	fmt.Println("TODO")
	for _, item := range todoList.Items {
		fmt.Printf("%s\n", item.Text)
	}
	// TODO: Add rendering logic here
}