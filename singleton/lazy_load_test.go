package singleton

import "testing"
import "github.com/stretchr/testify/assert"

func TestLazyLoad(t *testing.T) {
	singleton := GetInstance()
	assert.NotNil(t, singleton)
}
