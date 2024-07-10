package handlers

import (
	"fmt"
	"net/http"

	"github.com/robsongomes/htmx-starter/types"
	"github.com/robsongomes/htmx-starter/views"
)

type TodoStore interface {
	AllTodos() []types.Todo
	CreateTodo(desc string) types.Todo
}

type TodoHandler struct {
	store TodoStore
}

func NewTodoHandler(store TodoStore) *TodoHandler {
	return &TodoHandler{store: store}
}

func (th *TodoHandler) Home(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, views.TodoPage(th.store.AllTodos()))
}

func (th *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) error {
	desc := r.FormValue("description")
	todo := th.store.CreateTodo(desc)
	fmt.Println(todo)
	return render(w, r, views.TodoPage(th.store.AllTodos()))
}
