package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/go-lang-base/internal/tracker"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool}
}

func (r *RepoPg) Create(ctx context.Context, item tracker.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO items(id, name) VALUES($1, $2)`,
		item.ID, item.Name,
	)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	return nil
}

func (r *RepoPg) List(ctx context.Context) ([]tracker.Item, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, name FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) FindByNameLike(ctx context.Context, name string) ([]tracker.Item, error) {
	rows, err := r.pool.Query(
		ctx,
		`SELECT id, name FROM items WHERE LOWER(name) LIKE '%' || LOWER($1) || '%'`,
		name,
	)
	if err != nil {
		return nil, fmt.Errorf("r.pool.Query: %w", err)
	}
	defer rows.Close()

	var items []tracker.Item
	for rows.Next() {
		var item tracker.Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepoPg) Get(ctx context.Context, id string) (tracker.Item, error) {
	var item tracker.Item
	err := r.pool.QueryRow(ctx, `SELECT id, name FROM items WHERE id = $1`, id).Scan(&item.ID, &item.Name)

	return item, err
}

func (r *RepoPg) DeleteById(ctx context.Context, id string) error {
	result, err := r.pool.Exec(ctx, `DELETE FROM items WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return tracker.ErrNotFound
	}

	return nil
}

func (r *RepoPg) UpdateItem(ctx context.Context, item tracker.Item) error {
	result, err := r.pool.Exec(ctx, `UPDATE items SET name = $1 WHERE id = $2`, item.Name, item.ID)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return tracker.ErrNotFound
	}

	return nil
}
