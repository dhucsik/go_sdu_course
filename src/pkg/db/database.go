package db

import "time"

type Item struct {
	Id          int
	Title       string
	Price       float64
	Category    Category
	Description string
}

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

type Buyer struct {
	User           User
	FullName       string
	BillingAddress string
}

type Seller struct {
	User            User
	FullName        string
	BusinessAddress string
	PassportId      string
}

type Category struct {
	Id    int
	Title string
}

type Order struct {
	Id         int
	Buyer      Buyer
	Seller     Seller
	TotalPrice float64
	Date       time.Time
}

type OrderedItem struct {
	Id       int
	Order    Order
	Item     Item
	Quantity int
	Price    float64
}

type Payment struct {
	Id            int
	Order         Order
	PaymentMethod string
	TransactionId string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

var Items = make([]Item, 0)
var Users = make([]User, 0)
var Categories = make([]Category, 0)
var Orders = make([]Order, 0)
var Buyers = make([]Buyer, 0)
var Sellers = make([]Seller, 0)
var OrderedItems = make([]OrderedItem, 0)
var Payments = make([]Payment, 0)
