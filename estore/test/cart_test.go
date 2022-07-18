package test

import (
	"github.com/iseoluwaYN/estore/cart"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThatItemsCanBeAddedToCart(t *testing.T) {
	shoppingCart := &cart.ShoppingCart{}
	assert.NotNil(t, shoppingCart)
	assert.Empty(t, shoppingCart.Items)
	assert.Equal(t, 0, len(shoppingCart.Items))
	product := cart.Product{
		Id:          1,
		Name:        "milo",
		Price:       80.00,
		Description: "hot chocolate",
		Category:    "BEVERAGE",
	}

	item := cart.Item{
		Product:  product,
		Quantity: 2,
	}

	//when
	shoppingCart.AddItemToCart(item)
	//assert that item was added to cart
	assert.Equal(t, 1, len(shoppingCart.Items))

}
