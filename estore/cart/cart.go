package cart

type ShoppingCart struct {
	Items []Item
}

func (shoppingCart *ShoppingCart) AddItemToCart(item Item) {
	shoppingCart.Items = append(shoppingCart.Items, item)
}
