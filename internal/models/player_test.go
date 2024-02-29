package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewEmail_WhenEmailIsValid(t *testing.T) {
	e, err := NewEmail("example@example.com")

	assert.Nil(t, err)
	assert.NotNil(t, e)
}

func Test_NewEmail_WhenEmailIsNotValid(t *testing.T) {
	e, err := NewEmail("example")

	assert.NotNil(t, err)
	assert.Equal(t, Email(""), e)
}
