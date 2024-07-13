package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
	assert.InDelta(t, 10.0, product.Price, 0.01)
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	product, err := NewProduct("Product 2", -10.0)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidPrice, err)
	assert.Nil(t, product)
}

func TestProduct_RequiredName(t *testing.T) {
	product, err := NewProduct("", 10.0)
	assert.Error(t, err)
	assert.Equal(t, ErrNameIsRequired, err)
	assert.Nil(t, product)
}

func TestProduct_RequiredPrice(t *testing.T) {
	product, err := NewProduct("Product", 0.0)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_InvalidPrice(t *testing.T) {
	product, err := NewProduct("Product", -10.0)
	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}
