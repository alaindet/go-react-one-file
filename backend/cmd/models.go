package main

import (
	"fmt"
	"time"
)

var mockTodos []Todo = []Todo{
	{ID: "1", Text: "Buy Bread"},
	{ID: "2", Text: "Take a walk"},
	{ID: "3", Text: "Develop to-do apps"},
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

type CreateTodoDto struct {
	Text string `json:"text"`
}

type UpdateTodoDto struct {
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}

func RandomTodoID() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
