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
	tracker.AddItem(Item{id, name})
}

type GetUseCase struct{}

func (u GetUseCase) Done(_ Input, out Output, tracker *Tracker) {
	if len(tracker.Items) == 0 {
		out.Out("элементы не найдены")
	}

	for _, item := range tracker.Items {
		out.Out(item.toString())
	}
}

type DeleteUseCase struct{}

func (u DeleteUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("введите id для удаления:")
	id := in.Get()
	index := tracker.FindIndexById(id)

	if index == -1 {
		out.Out("элемент не найден")
		return
	}

	deleted := tracker.Items[index]
	tracker.Items = append(tracker.Items[:index], tracker.Items[index+1:]...)
	out.Out(fmt.Sprintf("удален элемент: %s", deleted.toString()))
}

type FindUseCase struct{}

func (u FindUseCase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("введите имя или часть имени:")
	name := strings.ToLower(in.Get())
	found := false

	for _, item := range tracker.Items {
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
	index := tracker.FindIndexById(id)

	if index == -1 {
		out.Out("элемент не найден")
		return
	}

	out.Out("введите новое название:")
	tracker.Items[index].Name = in.Get()
	out.Out(fmt.Sprintf("элемент обновлен: %s", tracker.Items[index].toString()))
}
