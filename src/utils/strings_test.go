package utils_test

import (
	"testing"

	"github.com/kamaal111/pocket-slate-api/src/utils"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input          string
	ExpectedOutput string
}

func TestPascalToSnakeCase(t *testing.T) {
	testCases := []testCase{
		{Input: "TargetLocale", ExpectedOutput: "target_locale"},
		{Input: "Target", ExpectedOutput: "target"},
	}

	for _, singleTestCase := range testCases {
		assert.Equal(t, utils.PascalToSnakeCase(singleTestCase.Input), singleTestCase.ExpectedOutput)
	}
}
