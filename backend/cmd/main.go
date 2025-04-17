package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"
)

//go:embed dist/*
var frontendEmbedded embed.FS

func main() {
	// Read input
	cfg := ReadServerConfig()

	// Create store based on input
	todosStore, err := CreateStoreFromServerConfig(cfg)
	if err != nil {
		panic(err)
	}

	server := http.NewServeMux()

	// Serve the frontend
	frontendFs, err := fs.Sub(frontendEmbedded, "dist")
	if err != nil {
		panic(err)
	}
	server.Handle("/", http.FileServer(http.FS(frontendFs)))

	// Register routes
	server.HandleFunc("POST /api/todos", createTodo(todosStore))
	server.HandleFunc("GET /api/todos", getTodos(todosStore))
	server.HandleFunc("GET /api/todos/{todoId}", getTodo(todosStore))
	server.HandleFunc("PUT /api/todos/{todoId}", updateTodo(todosStore))
	server.HandleFunc("DELETE /api/todos/{todoId}", deleteTodo(todosStore))

	// Register middleware
	serverWithMiddleware := corsMiddleware(server)
	serverWithMiddleware = loggingMiddleware(serverWithMiddleware)

	// Open new browser window when ready
	go func() {
		<-time.After(1 * time.Second)
		url := fmt.Sprintf("http://localhost:%s?apiPort=%s", cfg.Port, cfg.Port)
		err := OpenURLInBrowser(url)
		if err != nil {
			panic(err)
		}
	}()

	// Bootstrap the web server
	fmt.Printf("YATA - Yet Another Todo App has started on port %s\n", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, serverWithMiddleware)
	if err != nil {
		panic(err)
	}
}

func CreateStoreFromServerConfig(cfg ServerConfig) (*TodosStore, error) {
	if cfg.JSONDatabasePath != "" {
		return NewFilesystemTodosStore(cfg.JSONDatabasePath)
	}

	return NewInMemoryTodosStore(mockTodos)
}
