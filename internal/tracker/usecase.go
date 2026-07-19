package tracker

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(text string)
}

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
	DeleteById(ctx context.Context, id string) error
	UpdateItem(ctx context.Context, item Item) error
	FindByNameLike(ctx context.Context, name string) ([]Item, error)
}

type UseCase interface {
	Done(ctx context.Context, in Input, out Output, store Store) error
}

type AddUseCase struct{}

func (u AddUseCase) Done(ctx context.Context, in Input, out Output, store Store) error {
	out.Out("введите имя:")
	name := in.Get()
	id := uuid.New().String()
	item := Item{id, name}
	err := store.Create(ctx, item)
	if err != nil {
		return fmt.Errorf("ошибка при создании элемента: %w", err)
	}

	out.Out(fmt.Sprintf("новый элемент добавлен: %s", item.toString()))
	return nil
}

type GetUseCase struct{}

func (u GetUseCase) Done(ctx context.Context, _ Input, out Output, store Store) error {
	items, err := store.List(ctx)
	if err != nil {
		return fmt.Errorf("ошибка получения элементов: %w", err)
	}

	for _, item := range items {
		out.Out(item.toString())
	}
	return nil
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(ctx context.Context, in Input, out Output, store Store) error {
	out.Out("введите id для удаления:")
	id := in.Get()
	err := store.DeleteById(ctx, id)

	if err != nil {
		return fmt.Errorf("ошибка при удалении: %w", err)
	}

	out.Out(fmt.Sprintf("удален элемент c id: %s", id))
	return nil
}

type FindUseCase struct{}

func (u FindUseCase) Done(ctx context.Context, in Input, out Output, store Store) error {
	out.Out("введите имя или часть имени:")
	name := in.Get()

	items, err := store.FindByNameLike(ctx, name)
	if err != nil {
		return fmt.Errorf("ошибка при поиске по подстроке: %w", err)
	}

	if len(items) == 0 {
		out.Out("элементы не найдены")
		return nil
	}

	for _, item := range items {
		out.Out(fmt.Sprintf("найден элемент: %s", item.toString()))
	}

	return nil
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(ctx context.Context, in Input, out Output, store Store) error {
	out.Out("введите id для обновления:")
	id := in.Get()

	out.Out("введите новое название:")
	newName := in.Get()

	updatedItem := Item{id, newName}
	err := store.UpdateItem(ctx, updatedItem)

	if err != nil {
		return fmt.Errorf("ошибка при обновлении: %w", err)
	}

	out.Out(fmt.Sprintf("элемент обновлен: %s", updatedItem.toString()))
	return nil
}
