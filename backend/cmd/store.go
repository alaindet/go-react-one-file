package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type TodosStore struct {
	todos []Todo
	file  string
}

var (
	ErrTodosStore        = errors.New("todos store")
	ErrTodoAlreadyExists = fmt.Errorf("%w: already exists", ErrTodosStore)
	ErrTodoNotFound      = fmt.Errorf("%w: not found", ErrTodosStore)
	ErrInvalidJSONFile   = fmt.Errorf("%w: invalid JSON file", ErrTodosStore)
)

func NewInMemoryTodosStore(initialTodos []Todo) (*TodosStore, error) {

	store := &TodosStore{
		todos: initialTodos,
	}

	return store, nil
}

func NewFilesystemTodosStore(jsonDbPath string) (*TodosStore, error) {

	store := &TodosStore{
		todos: make([]Todo, 0),
	}

	todos, err := store.LoadFromJSON(jsonDbPath)
	if err != nil {
		return store, err
	}

	store.todos = todos
	store.file = jsonDbPath
	return store, nil
}

func (s *TodosStore) LoadFromJSON(jsonDbPath string) ([]Todo, error) {

	var todos []Todo

	// If no file is found, initialize an empty slice
	if _, err := os.Stat(jsonDbPath); os.IsNotExist(err) {
		return todos, nil
	}

	// Try loading content from file
	jsonRawContent, err := os.ReadFile(jsonDbPath)
	if err != nil {
		return todos, ErrInvalidJSONFile
	}

	err = json.Unmarshal(jsonRawContent, &todos)
	if err != nil {
		return todos, ErrInvalidJSONFile
	}

	return todos, nil
}

func (s *TodosStore) ExistsID(id string) bool {
	for _, todo := range s.todos {
		if todo.ID == id {
			return true
		}
	}
	return false
}

func (s *TodosStore) ExistsText(text string) bool {
	for _, todo := range s.todos {
		if todo.Text == text {
			return true
		}
	}
	return false
}

func (s *TodosStore) GetAll() []Todo {
	// return slices.Clone(s.todos)
	return s.todos
}

func (s *TodosStore) GetByID(id string) (Todo, error) {
	for _, todo := range s.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return Todo{}, ErrTodoNotFound
}

func (s *TodosStore) GetByText(text string) (Todo, error) {

	for _, todo := range s.todos {
		if strings.EqualFold(text, todo.Text) {
			return todo, nil
		}
	}

	return Todo{}, ErrTodoNotFound
}

func (s *TodosStore) MustGetByID(id string) Todo {
	todo, err := s.GetByID(id)
	if err != nil {
		panic(err)
	}

	return todo
}

func (s *TodosStore) Add(text string) (Todo, error) {

	exists := s.ExistsText(text)
	if exists {
		return Todo{}, ErrTodoAlreadyExists
	}

	todo := Todo{
		ID:     RandomTodoID(),
		Text:   text,
		IsDone: false,
	}

	s.todos = append(s.todos, todo)
	s.Persist()

	return todo, nil
}

func (s *TodosStore) Update(id string, dto UpdateTodoDto) (Todo, error) {

	foundIndex := -1

	for i, todo := range s.todos {

		// Error: another todo with the same text already exists
		if strings.EqualFold(todo.Text, dto.Text) && todo.ID != id {
			return Todo{}, ErrTodoAlreadyExists
		}

		if todo.ID == id {
			foundIndex = i
		}
	}

	if foundIndex == -1 {
		return Todo{}, ErrTodoNotFound
	}

	s.todos[foundIndex].Text = dto.Text
	s.todos[foundIndex].IsDone = dto.IsDone

	s.Persist()
	return s.todos[foundIndex], nil
}

func (s *TodosStore) Delete(id string) (Todo, error) {

	existingTodo, err := s.GetByID(id)
	if err != nil {
		return Todo{}, err
	}

	newTodos := make([]Todo, 0, len(s.todos)-1)

	for _, todo := range s.todos {
		if todo.ID != existingTodo.ID {
			newTodos = append(newTodos, todo)
		}
	}

	s.todos = newTodos
	s.Persist()

	return existingTodo, nil
}

func (s *TodosStore) Persist() {

	if s.file == "" {
		return
	}

	jsonData, err := json.Marshal(s.todos)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(s.file, jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Persisted JSON database to filesystem")
}
