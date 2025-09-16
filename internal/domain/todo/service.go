package todo

import "context"

type Service interface {
	CreateTodo(ctx context.Context, t *Todo) (*Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*Todo, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateTodo(ctx context.Context, t *Todo) (*Todo, error) {
	return s.repo.Create(ctx, t)
}

func (s *service) GetTodoByID(ctx context.Context, id int64) (*Todo, error) {
	return s.repo.GetByID(ctx, id)
}
