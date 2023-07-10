package utils_test

import (
	"errors"
	"testing"

	"github.com/kamaal111/pocket-slate-api/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestUnwrapEnvironment(t *testing.T) {
	t.Setenv("KEY", "secret")

	value, err := utils.UnwrapEnvironment("KEY")

	assert.Nil(t, err)
	assert.Equal(t, value, "secret")
}

func TestUnwrapEnvironmentNotFound(t *testing.T) {
	value, err := utils.UnwrapEnvironment("KEY")

	assert.Equal(t, err, errors.New("'KEY' not defined in the environment"))
	assert.Empty(t, value)
}
