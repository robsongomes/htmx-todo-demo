package store

import "github.com/robsongomes/htmx-starter/types"

type InMemoryStore struct {
	lastId int
	todos  []types.Todo
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		lastId: 2,
		todos: []types.Todo{
			{Id: 1, Description: "Dar banho no cachorro", Done: false},
			{Id: 2, Description: "Comprar ração", Done: true},
		},
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
