package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	res, err := New()

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
