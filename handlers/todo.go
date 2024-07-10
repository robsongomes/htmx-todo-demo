package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/robsongomes/htmx-starter/types"
	"github.com/robsongomes/htmx-starter/views"
)

type TodoStore interface {
	AllTodos() []types.Todo
	CreateTodo(desc string) types.Todo
	ToggleTodo(id int) types.Todo
	DeleteTodo(id int)
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
	return render(w, r, views.Todo(todo))
}

func (th *TodoHandler) ToggleTodo(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(r.PathValue("id"))
	todo := th.store.ToggleTodo(id)
	return render(w, r, views.Todo(todo))
}

func (th *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(r.PathValue("id"))
	th.store.DeleteTodo(id)
	w.Write([]byte(""))
	return nil
}
