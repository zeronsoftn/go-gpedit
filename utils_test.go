package go_gpedit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSysError_SuccessValue(t *testing.T) {
	result := WinErrHandler(nil, 0)
	assert.Equal(t, result, nil)
}

func TestSysError_FailureValue(t *testing.T) {
	result := WinErrHandler(nil, 0x81234567)
	assert.EqualErrorf(t, result, "System Error 0x81234567", "")
}
