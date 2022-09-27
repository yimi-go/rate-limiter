package rate_limiter

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrLimitExceed(t *testing.T) {
	err := ErrLimitExceed()
	assert.NotNil(t, err)
	assert.Same(t, errLimitExceed, err)
}

func TestIsErrLimitExceed(t *testing.T) {
	err := ErrLimitExceed()
	assert.True(t, IsErrLimitExceed(err))
	err = fmt.Errorf("wrapped: %w", err)
	assert.True(t, IsErrLimitExceed(err))
	assert.False(t, IsErrLimitExceed(io.EOF))
}
