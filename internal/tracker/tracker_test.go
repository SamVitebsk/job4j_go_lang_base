package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTracker(t *testing.T) {
	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.UpdateItem(item)
		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("update existing item", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := tracker.AddItem(item)
		require.NoError(t, err)

		updatedItem := Item{
			ID:   "1",
			Name: "Updated Item",
		}
		err = tracker.UpdateItem(updatedItem)

		assert.NoError(t, err)
		assert.Equal(t, updatedItem, tracker.items[0])
	})

	t.Run("add item success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := tracker.AddItem(item)
		assert.NoError(t, err)
		assert.Equal(t, item, tracker.items[0])
	})

	t.Run("error add - item already exists", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		firstItem := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := tracker.AddItem(firstItem)
		require.NoError(t, err)

		secondItem := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err2 := tracker.AddItem(secondItem)

		assert.Equal(t, firstItem, tracker.items[0])
		assert.ErrorIs(t, err2, ErrItemAlreadyExists)
	})

	t.Run("get success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		firstItem := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := tracker.AddItem(firstItem)
		require.NoError(t, err)

		secondItem := Item{
			ID:   "2",
			Name: "First Item",
		}
		_, err2 := tracker.AddItem(secondItem)
		require.NoError(t, err2)

		actualItems := tracker.GetItems()

		assert.Equal(t, firstItem, actualItems[0])
		assert.Equal(t, secondItem, actualItems[1])
	})

	t.Run("error delete - item not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		firstItem := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.DeleteItem(firstItem.ID)

		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("delete success", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		firstItem := Item{
			ID:   "1",
			Name: "First Item",
		}
		_, err := tracker.AddItem(firstItem)
		require.NoError(t, err)

		err2 := tracker.DeleteItem(firstItem.ID)

		assert.NoError(t, err2)
		assert.Equal(t, 0, len(tracker.items))
	})
}
