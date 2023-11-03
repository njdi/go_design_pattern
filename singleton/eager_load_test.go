package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEagerLoad(t *testing.T) {
	singleton := GetEagerInstance()
	assert.NotNil(t, singleton)
}
