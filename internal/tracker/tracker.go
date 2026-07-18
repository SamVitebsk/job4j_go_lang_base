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
	Items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.Items = append(t.Items, item)
}

func (t *Tracker) GetItems() []Item {
	return t.Items
}

func (t *Tracker) FindIndexById(id string) int {
	for i, item := range t.Items {
		if item.ID == id {
			return i
		}
	}

	return -1
}
