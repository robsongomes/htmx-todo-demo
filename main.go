package main

import (
	"net/http"

	"github.com/robsongomes/htmx-starter/handlers"
	"github.com/robsongomes/htmx-starter/store"
)

func main() {

	store := store.NewInMemoryStore()
	todoHandler := handlers.NewTodoHandler(store)

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("GET /", handlers.HTTPHandler(handlers.HomeHandler))

	http.Handle("GET /todos", handlers.HTTPHandler(todoHandler.Home))
	http.Handle("GET /todos/filter", handlers.HTTPHandler(todoHandler.FilterTodos))
	http.Handle("POST /todos", handlers.HTTPHandler(todoHandler.CreateTodo))
	http.Handle("PUT /todos/{id}", handlers.HTTPHandler(todoHandler.ToggleTodo))
	http.Handle("DELETE /todos/{id}", handlers.HTTPHandler(todoHandler.DeleteTodo))

	http.ListenAndServe(":3000", nil)
}
