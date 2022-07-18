package user

import (
	"github.com/iseoluwaYN/estore/cart"
	"time"
)

type CardType string

type Customer struct {
	Age                int
	Email              string
	PhoneNumber        string
	Name               string
	Password           string
	Address            Address
	BillingInformation []BillingInformation
	ShoppingCart       cart.ShoppingCart
}

type CreditCardInfo struct {
	Cvv              int
	ExpirationYear   int
	Month            time.Month
	CreditCardNumber string
	Name             string
	CardType         CardType
}

type BillingInformation struct {
	ReceiversNumber string
	ReceiversName   string
	DeliveryAddress Address
	CreditCardInfo  CreditCardInfo
}
