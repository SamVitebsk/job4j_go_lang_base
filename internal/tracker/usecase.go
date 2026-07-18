package tracker

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UseCase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUseCase struct{}

func (u AddUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("введите имя:")
	name := in.Get()
	id := uuid.New().String()
	item, err := tracker.AddItem(Item{id, name})
	if err != nil {
		out.Out(fmt.Sprintf("элемент с id = %s уже существует", id))
		return
	}

	out.Out(fmt.Sprintf("новый элемент добавлен: %s", item.toString()))
}

type GetUseCase struct{}

func (u GetUseCase) Done(_ Input, out Output, tracker *Tracker) {
	if len(tracker.items) == 0 {
		out.Out("элементы не найдены")
		return
	}

	for _, item := range tracker.items {
		out.Out(item.toString())
	}
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("введите id для удаления:")
	id := in.Get()
	err := tracker.DeleteItem(id)

	if err != nil {
		out.Out("элемент не найден")
		return
	}

	out.Out(fmt.Sprintf("удален элемент c id: %s", id))
}

type FindUseCase struct{}

func (u FindUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("введите имя или часть имени:")
	name := strings.ToLower(in.Get())
	found := false

	for _, item := range tracker.items {
		if strings.Contains(strings.ToLower(item.Name), name) {
			out.Out(fmt.Sprintf("найден элемент: %s", item.toString()))
			found = true
		}
	}

	if !found {
		out.Out("элементы не найдены")
	}
}

type UpdateUseCase struct{}

func (u UpdateUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("введите id для обновления:")
	id := in.Get()

	out.Out("введите новое название:")
	newName := in.Get()

	updatedItem := Item{id, newName}
	err := tracker.UpdateItem(updatedItem)

	if err != nil {
		out.Out("элемент не найден")
		return
	}

	out.Out(fmt.Sprintf("элемент обновлен: %s", updatedItem.toString()))
}
