package go_gpedit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/windows"
	"testing"
)

func TestSysError_UseParent(t *testing.T) {
	parent := fmt.Errorf("HELLO WORLD")
	result := WinErrHandler(parent, 0)
	assert.Equal(t, result, parent)
}

func TestSysError_SuccessValue(t *testing.T) {
	result := WinErrHandler(nil, 0)
	assert.Equal(t, result, nil)
}

func TestSysError_FailureValue(t *testing.T) {
	result := WinErrHandler(nil, 12345678)
	expected := windows.Errno(12345678)
	assert.EqualError(t, result, expected.Error())
	code, ok := result.(windows.Errno)
	assert.True(t, ok)
	assert.Equal(t, int(code), 12345678)
}
