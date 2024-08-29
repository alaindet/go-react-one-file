package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateTodosStore(t *testing.T) {

	initialTodos := []Todo{
		{ID: "1", Text: "Foo"},
		{ID: "2", Text: "Bar"},
	}

	t.Run("Creates without options", func(t *testing.T) {
		s := NewTodosStore()
		todos := s.GetAll()
		emptyTodos := make([]Todo, 0)
		if !reflect.DeepEqual(todos, emptyTodos) {
			t.Errorf("Initial todos slice is not empty")
		}
	})

	t.Run("Creates with initial slice", func(t *testing.T) {
		s := NewTodosStore(WithTodos(initialTodos))
		todos := s.GetAll()
		if !reflect.DeepEqual(todos, initialTodos) {
			t.Errorf("Initial todos are not set")
		}
	})
}

func TestReadTodosStore(t *testing.T) {
	initialTodos := []Todo{
		{ID: "1", Text: "Foo"},
		{ID: "2", Text: "Bar"},
	}

	s := NewTodosStore(WithTodos(initialTodos))

	t.Run("Checks for existing ID", func(t *testing.T) {
		existingID := "1"
		exists := s.ExistsID(existingID)
		if !exists {
			t.Errorf("ID %q should exist in the store", existingID)
		}
	})

	t.Run("Checks for non-existing ID", func(t *testing.T) {
		nonExistingID := "42"
		exists := s.ExistsID(nonExistingID)
		if exists {
			t.Errorf("ID %q should not exist in the store", nonExistingID)
		}
	})

	t.Run("Checks for existing text", func(t *testing.T) {
		existingText := "Foo"
		exists := s.ExistsText(existingText)
		if !exists {
			t.Errorf("Text %q should exist in the store", existingText)
		}
	})

	t.Run("Checks for non-existing ID", func(t *testing.T) {
		nonExistingText := "nope"
		exists := s.ExistsText(nonExistingText)
		if exists {
			t.Errorf("Text %q should not exist in the store", nonExistingText)
		}
	})

	t.Run("Gets a todo by searching for existing ID", func(t *testing.T) {
		expected := Todo{ID: "1", Text: "Foo"}
		todo, err := s.GetByID(expected.ID)
		assertNoError(t, err)
		assertTodosEqual(t, todo, expected)
	})

	t.Run("Gets an error by searching for non-existing ID", func(t *testing.T) {
		todo, err := s.GetByID("42")
		assertTodoIsEmpty(t, todo)
		assertError(t, err)
	})

	t.Run("Gets a panic by searching for non-existing ID", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic, none given")
			}
		}()
		_ = s.MustGetByID("42")
	})

	t.Run("Gets a todo by searching for existing text (case exact)", func(t *testing.T) {
		expected := Todo{ID: "1", Text: "Foo"}
		todo, err := s.GetByText(expected.Text)
		assertNoError(t, err)
		assertTodosEqual(t, todo, expected)
	})

	t.Run("Gets a todo by searching for existing text (case insensitive)", func(t *testing.T) {
		textQuery := "fOo"
		expected := Todo{ID: "1", Text: "Foo"}
		todo, err := s.GetByText(textQuery)
		assertNoError(t, err)
		assertTodosEqual(t, todo, expected)
	})

	t.Run("Gets an error by searching for non-existing text", func(t *testing.T) {
		todo, err := s.GetByText("nope")
		assertTodoIsEmpty(t, todo)
		assertError(t, err)
	})
}

func TestWriteTodosStore(t *testing.T) {

	setupStore := func() *TodosStore {
		return NewTodosStore(WithTodos([]Todo{
			{ID: "1", Text: "Foo"},
			{ID: "2", Text: "Bar"},
		}))
	}

	t.Run("Updates a todo", func(t *testing.T) {
		s := setupStore()
		expected := Todo{ID: "1", Text: "Foo updated", IsDone: true}
		dto := UpdateTodoDto{Text: expected.Text, IsDone: expected.IsDone}
		updatedTodo, err := s.Update(expected.ID, dto)
		assertNoError(t, err)
		assertTodosEqual(t, updatedTodo, expected)
		result, err := s.GetByID(expected.ID)
		assertNoError(t, err)
		assertTodosEqual(t, result, expected)
	})

	t.Run("Fails to update a todo with non-existing ID", func(t *testing.T) {
		s := setupStore()
		expected := Todo{ID: "42", Text: "Foo updated", IsDone: true}
		dto := UpdateTodoDto{Text: expected.Text, IsDone: expected.IsDone}
		_, err := s.Update(expected.ID, dto)
		assertErrorIs(t, err, ErrTodoNotFound)
	})

	t.Run("Fails to update a todo with existing text", func(t *testing.T) {
		s := setupStore()
		expected := Todo{ID: "1", Text: "Bar", IsDone: false}
		dto := UpdateTodoDto{Text: expected.Text, IsDone: expected.IsDone}
		_, err := s.Update(expected.ID, dto)
		assertErrorIs(t, err, ErrTodoAlreadyExists)
	})

	t.Run("Updates a todo with itself", func(t *testing.T) {
		s := setupStore()
		todo := Todo{ID: "1", Text: "Foo", IsDone: false}
		dto := UpdateTodoDto{Text: todo.Text, IsDone: todo.IsDone}
		updatedTodo, err := s.Update(todo.ID, dto)
		assertNoError(t, err)
		assertTodosEqual(t, updatedTodo, todo)
	})

	t.Run("Deletes a todo", func(t *testing.T) {
		s := setupStore()
		countBefore := len(s.GetAll())
		todo := Todo{ID: "1", Text: "Foo", IsDone: false}
		deletedTodo, err := s.Delete("1")
		countAfter := len(s.GetAll())
		assertNoError(t, err)
		assertTodosEqual(t, deletedTodo, todo)
		assertEqual(t, countAfter, countBefore-1)
	})

	t.Run("Fails to delete a todo with non-existing ID", func(t *testing.T) {
		s := setupStore()
		countBefore := len(s.GetAll())
		_, err := s.Delete("42")
		countAfter := len(s.GetAll())
		assertErrorIs(t, err, ErrTodoNotFound)
		assertEqual(t, countAfter, countBefore)
	})
}

func assertEqual[T any](t *testing.T, given T, expected T) {
	t.Helper()
	if !reflect.DeepEqual(given, expected) {
		t.Errorf("Expected %#v, given %#v", expected, given)
	}
}

func assertNoError(t *testing.T, givenErr error) {
	t.Helper()
	if givenErr != nil {
		t.Errorf("Expected nil error, given %q", givenErr.Error())
	}
}

func assertError(t *testing.T, givenErr error) {
	t.Helper()
	if givenErr == nil {
		t.Errorf("Expected an error, given nil")
	}
}

func assertErrorIs(t *testing.T, given error, expected error) {
	t.Helper()
	if !errors.Is(given, expected) {
		t.Errorf("Expected an error, given nil")
	}
}

func assertTodosEqual(t *testing.T, given, expected Todo) {
	t.Helper()
	sameID := given.ID == expected.ID
	sameText := given.Text == expected.Text
	sameIsDone := given.IsDone == expected.IsDone
	if !sameID || !sameText || !sameIsDone {
		t.Errorf("Todos are not equal. Expected %#v, given %#v", expected, given)
	}
}

func assertTodoIsEmpty(t *testing.T, given Todo) {
	t.Helper()
	if given.ID != "" || given.Text != "" || given.IsDone != false {
		t.Errorf("Expected empty todo, given %#v", given)
	}
}
