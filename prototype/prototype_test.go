package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrototype(t *testing.T) {
	prototype := &ConcretePrototype{
		Field1: "field1",
		Field2: 1,
	}

	clone := prototype.Clone()

	assert.Equal(t, prototype.Field1, clone.Field1)
	assert.Equal(t, prototype.Field2, clone.Field2)
}
