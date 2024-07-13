package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/robsongomes/htmx-starter/services"
	"github.com/robsongomes/htmx-starter/types"
	"github.com/robsongomes/htmx-starter/views"
)

type TodoStore interface {
	AllTodos() []types.Todo
	CreateTodo(desc string) types.Todo
	ToggleTodo(id int) types.Todo
	DeleteTodo(id int)
	Filter(expr string) []types.Todo
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

func (th *TodoHandler) FilterTodos(w http.ResponseWriter, r *http.Request) error {
	filter := r.URL.Query().Get("filter")
	return render(w, r, views.TodoList(th.store.Filter(filter)))
}

func (th *TodoHandler) ValidateTodoDescription(w http.ResponseWriter, r *http.Request) error {
	description := r.FormValue("description")
	res, err := services.Validate(description)
	if err != nil {
		return err
	}
	if res.BadWordsTotal > 0 {
		msg := "Palavras impróprias: "
		msg += res.BadWords()
		return render(w, r, views.TodoValidationMessage(msg))
	}
	return render(w, r, views.TodoValidationMessage(""))
}
