package todo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/internal/cache"
	"time"
)

type Repository interface {
	Create(ctx context.Context, t *Todo) (*Todo, error)
	GetByID(ctx context.Context, id int64) (*Todo, error)
}

type repository struct {
	db    *sql.DB
	cache cache.Cache
}

func NewRepository(db *sql.DB, cache cache.Cache) Repository {
	return &repository{db: db, cache: cache}
}

func (r *repository) Create(ctx context.Context, t *Todo) (*Todo, error) {
	query := "INSERT INTO todos (title, completed) VALUES (?, ?) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, t.Title, t.Completed).Scan(&t.ID)
	if err != nil {
		return nil, err
	}

	if r.cache != nil {
		data, _ := json.Marshal(t)
		_ = r.cache.Set(fmt.Sprintf("todo:%d", t.ID), string(data), 10*time.Minute)
	}

	return t, nil
}

func (r *repository) GetByID(ctx context.Context, id int64) (*Todo, error) {
	key := fmt.Sprintf("todo:%d", id)
	if r.cache != nil {
		if val, err := r.cache.Get(key); err == nil {
			var t Todo
			if err := json.Unmarshal([]byte(val), &t); err == nil {
				return &t, nil
			}
		}
	}

	var t Todo
	query := "SELECT id, title, completed FROM todos WHERE id=?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		return nil, err
	}

	if r.cache != nil {
		data, _ := json.Marshal(t)
		_ = r.cache.Set(key, string(data), 10*time.Minute)
	}

	return &t, nil
}
