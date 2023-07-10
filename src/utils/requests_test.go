package utils_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kamaal111/pocket-slate-api/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestUnwrapURLQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)

	value, err := utils.UnwrapURLQuery(request, "word")

	assert.Nil(t, err)
	assert.Equal(t, value, "abc")
}

func TestUnwrapURLQueryNotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)

	value, err := utils.UnwrapURLQuery(request, "something")

	assert.Equal(t, err, errors.New("'something' is not defined in the query params"))
	assert.Empty(t, value)
}
