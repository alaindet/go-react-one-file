package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RequestHandler func(http.ResponseWriter, *http.Request)

func createTodo(store *TodosStore) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		// Parse and validate
		var dto CreateTodoDto
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todo, err := store.Add(dto.Text)
		if err != nil {
			message := fmt.Sprintf("Text %q already exists", dto.Text)
			SendJson(w, http.StatusConflict, Resp(message, nil))
			return
		}

		message := fmt.Sprintf("Todo %q created", dto.Text)
		SendJson(w, http.StatusCreated, Resp(message, todo))
	}
}

func getTodos(store *TodosStore) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		message := "Get all todos"
		todos := store.GetAll()
		SendJson(w, http.StatusOK, Resp(message, todos))
	}
}

func getTodo(store *TodosStore) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		todoId := r.PathValue("todoId")

		todo, err := store.GetByID(todoId)
		if err != nil {
			message := fmt.Sprintf("Todo with ID %q not found", todoId)
			SendJson(w, http.StatusNotFound, Resp(message, nil))
			return
		}

		message := fmt.Sprintf("Get todo with ID %q", todoId)
		SendJson(w, http.StatusOK, Resp(message, todo))
	}
}

func updateTodo(store *TodosStore) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		todoId := r.PathValue("todoId")

		// Parse and validate
		var dto UpdateTodoDto
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedTodo, err := store.Update(todoId, dto)
		if err != nil {
			if errors.Is(err, ErrTodoNotFound) {
				message := fmt.Sprintf("Todo with ID %q not found", todoId)
				SendJson(w, http.StatusNotFound, Resp(message, nil))
				return
			}

			if errors.Is(err, ErrTodoAlreadyExists) {
				message := fmt.Sprintf("A todo with text %q already exists", dto.Text)
				SendJson(w, http.StatusNotFound, Resp(message, nil))
				return
			}
		}

		message := fmt.Sprintf("Todo with ID %q was updated", todoId)
		SendJson(w, http.StatusOK, Resp(message, updatedTodo))
	}
}

func deleteTodo(store *TodosStore) RequestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		todoId := r.PathValue("todoId")

		deletedTodo, err := store.Delete(todoId)
		if err != nil {
			message := fmt.Sprintf("Todo with ID %q not found", todoId)
			SendJson(w, http.StatusNotFound, Resp(message, nil))
			return
		}

		message := fmt.Sprintf("Todo with ID %q was deleted", todoId)
		SendJson(w, http.StatusOK, Resp(message, deletedTodo))
	}
}
