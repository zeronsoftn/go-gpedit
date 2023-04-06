package go_gpedit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	result := WinErrHandler(nil, 0x81234567)
	assert.EqualError(t, result, "System Error 0x81234567")
}
