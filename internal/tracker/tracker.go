package tracker

import (
	"fmt"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	_, ok := t.indexOf(item.ID)
	if ok {
		return item, ErrItemAlreadyExists
	}

	t.items = append(t.items, item)

	return item, nil
}

func (t *Tracker) GetItems() []Item {
	return append([]Item(nil), t.items...)
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.items {
		if item.ID == id {
			return i, true
		}
	}

	return -1, false
}

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}

	t.items[index] = item
	return nil
}

func (t *Tracker) DeleteItem(id string) error {
	index, ok := t.indexOf(id)

	if !ok {
		return ErrNotFound
	}

	t.items = append(t.items[:index], t.items[index+1:]...)
	return nil
}
