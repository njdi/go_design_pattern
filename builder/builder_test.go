package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {
	product := NewProductBuilder().
		WithName("Product 1").
		WithCategory("Category 1").
		WithPrice(9.99).
		Build()

	assert.Equal(t, "Product 1", product.name)
	assert.Equal(t, "Category 1", product.category)
	assert.Equal(t, 9.99, product.price)
}
