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
			{Id: 1, Description: "Dar banho no cachorro"},
			{Id: 2, Description: "Comprar ração"},
		},
	}
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
