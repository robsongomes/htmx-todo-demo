package store

import (
	"strings"

	"github.com/robsongomes/htmx-starter/types"
)

type InMemoryStore struct {
	lastId int
	todos  []types.Todo
}

func NewInMemoryStore() *InMemoryStore {
	todos := []types.Todo{
		{Id: 1, Description: "Dar banho no cachorro", Done: false},
		{Id: 2, Description: "Comprar ração", Done: true},
	}
	return &InMemoryStore{
		lastId: 2,
		todos:  todos,
	}
}

func (ms *InMemoryStore) indexOf(id int) int {
	for idx, todo := range ms.todos {
		if todo.Id == id {
			return idx
		}
	}
	return -1
}

func (ms *InMemoryStore) AllTodos() []types.Todo {
	return ms.todos
}

func (ms *InMemoryStore) CreateTodo(desc string) types.Todo {
	ms.lastId++
	todo := types.Todo{Id: ms.lastId, Description: desc, Done: false}
	ms.todos = append(ms.todos, todo)
	return todo
}

func (ms *InMemoryStore) ToggleTodo(id int) types.Todo {
	idx := ms.indexOf(id)
	ms.todos[idx].Done = !ms.todos[idx].Done
	return ms.todos[idx]
}

func (ms *InMemoryStore) DeleteTodo(id int) {
	idx := ms.indexOf(id)
	ms.todos = append(ms.todos[:idx], ms.todos[idx+1:]...)
}

func (ms *InMemoryStore) Filter(expr string) (res []types.Todo) {
	for _, t := range ms.todos {
		if strings.Contains(strings.ToLower(t.Description), strings.ToLower(expr)) {
			res = append(res, t)
		}
	}
	return
}
