package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate_NilRequest_ReturnError(t *testing.T) {

	rsl := base.Validate(nil)

	expected := []string{"req is nil"}

	assert.Equal(t, rsl, expected)
}

func Test_Validate_EmptyUserID_ReturnUserIDError(t *testing.T) {

	rsl := base.Validate(&base.ValidateRequest{
		UserID:      "",
		Title:       "test title",
		Description: "test description",
	})

	expected := []string{"UserID is required"}

	assert.Equal(t, rsl, expected)
}

func Test_Validate_EmptyTitle_ReturnTitleError(t *testing.T) {

	rsl := base.Validate(&base.ValidateRequest{
		UserID:      "123",
		Title:       "",
		Description: "test description",
	})

	expected := []string{"Title is required"}

	assert.Equal(t, rsl, expected)
}

func Test_Validate_EmptyDescription_ReturnDescriptionError(t *testing.T) {

	rsl := base.Validate(&base.ValidateRequest{
		UserID:      "123",
		Title:       "test title",
		Description: "",
	})

	expected := []string{"Description is required"}

	assert.Equal(t, rsl, expected)
}

func Test_Validate_AllFieldsEmpty_ReturnAllFieldsError(t *testing.T) {

	rsl := base.Validate(&base.ValidateRequest{
		UserID:      "",
		Title:       "",
		Description: "",
	})

	expected := []string{
		"UserID is required",
		"Title is required",
		"Description is required",
	}

	assert.Equal(t, rsl, expected)
}

func Test_Validate_AllFieldsPresent_ReturnEmptyError(t *testing.T) {

	rsl := base.Validate(&base.ValidateRequest{
		UserID:      "123",
		Title:       "test title",
		Description: "test description",
	})

	expected := make([]string, 0)

	assert.Equal(t, rsl, expected)
}
