package user

type Address struct {
	CityName    string
	Country     string
	HouseNumber int
	Street      string
	State       string
}

type Seller struct {
	Age         int
	Email       string
	PhoneNumber string
	Address     Address
	Name        string
	Password    string
}
