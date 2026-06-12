package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Add(t *testing.T) {

	rsl := base.Add(1, 2)
	expected := 3

	assert.Equal(t, rsl, expected)

}
