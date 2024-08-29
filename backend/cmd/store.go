package main

import (
	"errors"
	"strings"
)

type TodosStore struct {
	todos []Todo
}

var (
	ErrTodoAlreadyExists = errors.New("todo already exists")
	ErrTodoNotFound      = errors.New("todo not found")
)

type TodosStoreOption func(*TodosStore)

func NewTodosStore(options ...TodosStoreOption) *TodosStore {
	store := &TodosStore{
		todos: make([]Todo, 0),
	}

	for _, option := range options {
		option(store)
	}

	return store
}

func WithTodos(initialTodos []Todo) TodosStoreOption {
	return func(s *TodosStore) {
		s.todos = initialTodos
	}
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
	return todo, nil
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
	textQuery := strings.ToLower(text)

	for _, todo := range s.todos {
		if strings.ToLower(todo.Text) == textQuery {
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

func (s *TodosStore) Update(id string, dto UpdateTodoDto) (Todo, error) {

	existingTodo, err := s.GetByID(id)
	if err != nil {
		return Todo{}, err
	}

	existsByText, err := s.GetByText(dto.Text)
	if err == nil && existsByText.ID != existingTodo.ID {
		return Todo{}, ErrTodoAlreadyExists
	}

	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos[i].Text = dto.Text
			s.todos[i].IsDone = dto.IsDone
			return s.todos[i], nil
		}
	}

	return Todo{}, ErrTodoNotFound
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

	return existingTodo, nil
}
