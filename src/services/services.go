package services

import (
	"fmt"
	"module/database"
	"time"
)

func CreateUser() database.User {
	var username string
	var id int
	var email string
	var password string

	if len(database.Users) > 0 {
		id = database.Users[len(database.Users)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter username:")
	fmt.Scanf("%s\n", &username)
	fmt.Println("Enter email:")
	fmt.Scanf("%s\n", &email)
	fmt.Println("Enter password:")
	fmt.Scanf("%s\n", &password)

	user := database.User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}

	database.Users = append(database.Users, database.User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	})

	return user
}

func CreateBuyer() {
	user := CreateUser()
	var fullName string
	var billingAddress string

	fmt.Println("Enter your full name:")
	fmt.Scanf("%s\n", &fullName)
	fmt.Println("Enter your billing address:")
	fmt.Scanf("%s\n", billingAddress)

	database.Buyers = append(database.Buyers, database.Buyer{
		User:           user,
		FullName:       fullName,
		BillingAddress: billingAddress,
	})
}

func CreateSeller() {
	user := CreateUser()
	var fullName string
	var businessAddress string
	var passportId string

	fmt.Println("Enter your full name:")
	fmt.Scanf("%s\n", &fullName)
	fmt.Println("Enter your business address:")
	fmt.Scanf("%s\n", &businessAddress)
	fmt.Println("Enter your passport id:")
	fmt.Scanf("%s\n", &passportId)

	database.Sellers = append(database.Sellers, database.Seller{
		User:            user,
		FullName:        fullName,
		BusinessAddress: businessAddress,
		PassportId:      passportId,
	})
}

func CreateCategory() {
	var id int
	var category string

	if len(database.Categories) > 0 {
		id = database.Categories[len(database.Categories)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter category:")
	fmt.Scanf("%s\n", &category)

	database.Categories = append(database.Categories, database.Category{
		Id:    id,
		Title: category,
	})
}

func CreateItem() {
	var id int
	var title string
	var price float64
	var categoryId int
	var description string
	var category database.Category

	if len(database.Items) > 0 {
		id = database.Items[len(database.Items)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter title:")
	fmt.Scanf("%s,\n", &title)
	fmt.Println("Enter price:")
	fmt.Scanf("%f\n", &price)
	fmt.Println("Choose the category(enter the id):")
	ListCategories()
	fmt.Scanf("%d\n", &categoryId)
	fmt.Println("Enter description:")
	fmt.Scanf("%s\n", &description)

	for _, cat := range database.Categories {
		if cat.Id == categoryId {
			category = cat
			break
		}
	}

	database.Items = append(database.Items, database.Item{
		Id:          id,
		Title:       title,
		Price:       price,
		Category:    category,
		Description: description,
	})
}

func CreateOrder() {
	var id int
	var buyerId int
	var sellerId int
	date := time.Now()
	var totalPrice float64
	var buyer database.Buyer
	var seller database.Seller

	if len(database.Orders) > 0 {
		id = database.Orders[len(database.Orders)-1].Id + 1
	} else {
		id = 1
	}

	ListBuyers()
	fmt.Println("Enter id of the buyer:")
	fmt.Scanf("%d\n", &buyerId)
	ListSellers()
	fmt.Println("Enter id of the seller:")
	fmt.Scanf("%d\n", &sellerId)
	fmt.Println("Enter total price:")
	fmt.Scanf("%f\n", &totalPrice)

	for _, buy := range database.Buyers {
		if buy.User.Id == buyerId {
			buyer = buy
			break
		}
	}

	for _, sell := range database.Sellers {
		if sell.User.Id == sellerId {
			seller = sell
			break
		}
	}

	database.Orders = append(database.Orders, database.Order{
		Id:         id,
		Buyer:      buyer,
		Seller:     seller,
		TotalPrice: totalPrice,
		Date:       date,
	})
}

func ListUsers() {
	for _, user := range database.Users {
		fmt.Printf("Id: %d\t| Username: %s\t| Email: %s\t| Password: %s\n",
			user.Id, user.Username, user.Email, user.Password)
	}
}

func ListBuyers() {
	for _, buyer := range database.Buyers {
		fmt.Printf("Id: %d\t| Full name: %s\t| Billing address: %s\n",
			buyer.User.Id, buyer.FullName, buyer.BillingAddress)
	}
}

func ListSellers() {
	for _, seller := range database.Sellers {
		fmt.Printf("Id: %d\t| Full name: %s\t| Business address: %s\t| Passport id: %s\n",
			seller.User.Id, seller.FullName, seller.BusinessAddress, seller.PassportId)
	}
}

func ListCategories() {
	for _, category := range database.Categories {
		fmt.Printf("Id: %d\t| Title: %s\n", category.Id, category.Title)
	}
}

func ListItems() {
	for _, item := range database.Items {
		fmt.Printf("Id: %d\t| Title: %s\t| Price: %f\t| Category Id: %d\t| Description: %s\n",
			item.Id, item.Title, item.Price, item.Category.Id, item.Description)
	}
}

func ListOrders() {
	for _, order := range database.Orders {
		year, month, date := order.Date.Date()
		fmt.Printf("Id: %d\t| Buyer Id: %d\t| Seller Id: %d\t| Total price: %f\t| Date: %s\n",
			order.Id, order.Buyer.User.Id, order.Seller.User.Id, order.TotalPrice,
			fmt.Sprintf("%d %s, %d", date, month, year))
	}
}

func UpdateUser() {
	var id int
	var username string

	fmt.Println("Enter user id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter new username:")
	fmt.Scanf("%s\n", &username)

	for i, user := range database.Users {
		if user.Id == id {
			database.Users[i].Username = username
			break
		}
	}
}

func UpdateCategory() {
	var id int
	var categoryTitle string

	fmt.Println("Enter category id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter new category title:")
	fmt.Scanf("%s\n", &categoryTitle)

	for i, category := range database.Categories {
		if category.Id == id {
			database.Categories[i].Title = categoryTitle
			break
		}
	}
}

func UpdateItem() {
	var id int
	var itemTitle string
	var price float64
	var description string

	fmt.Println("Enter item id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter new item title:")
	fmt.Scanf("%s\n", &itemTitle)
	fmt.Println("Enter new item price:")
	fmt.Scanf("%f\n", &price)
	fmt.Println("Enter new description:")
	fmt.Scanf("%s\n", &description)

	for i, item := range database.Items {
		if item.Id == id {
			database.Items[i].Title = itemTitle
			database.Items[i].Price = price
			database.Items[i].Description = description
			break
		}
	}
}

func DeleteUser() {
	var id int

	fmt.Println("Enter user id:")
	fmt.Scanf("%d\n", id)

	for i, user := range database.Users {
		if user.Id == id {
			database.Users = append(database.Users[:i], database.Users[i+1:]...)
			break
		}
	}
}

func DeleteCategory() {
	var id int

	fmt.Println("Enter category id:")
	fmt.Scanf("%d\n", id)

	for i, category := range database.Categories {
		if category.Id == id {
			database.Categories = append(database.Categories[:i], database.Categories[i+1:]...)
			break
		}
	}
}

func DeleteItem() {
	var id int

	fmt.Println("Enter item id:")
	fmt.Scanf("%d\n", id)

	for i, item := range database.Items {
		if item.Id == id {
			database.Items = append(database.Items[:i], database.Items[i+1:]...)
			break
		}
	}
}

func DeleteOrder() {
	var id int

	fmt.Println("Enter order id:")
	fmt.Scanf("%d\n", id)

	for i, order := range database.Orders {
		if order.Id == id {
			database.Orders = append(database.Orders[:i], database.Orders[i+1:]...)
			break
		}
	}
}
