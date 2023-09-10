package resolver

import (
	"strconv"

	"github.com/tf63/go-graph-exp/api/graph"
	"github.com/tf63/go-graph-exp/internal/entity"
)

func TodoDTO(e *entity.Todo) graph.Todo {
	id := strconv.Itoa(int(e.Id))

	return graph.Todo{
		ID:   id,
		Text: e.Text,
		Done: e.Done,
		// CreatedAt: e.CreatedAt,
		// UpdatedAt: e.UpdatedAt,
	}
}

func NewTodoDTO(s *graph.NewTodo) entity.NewTodo {
	return entity.NewTodo{
		Text: s.Text,
	}
}
