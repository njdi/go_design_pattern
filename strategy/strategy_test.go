package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrategy(t *testing.T) {
	customer := &Context{&RegularCustomerStrategy{}}
	assert.Equal(t, customer.ExecuteStrategy(100), float32(100))

	customer = &Context{&PremiumCustomerStrategy{}}
	assert.Equal(t, customer.ExecuteStrategy(100), float32(90))
}
