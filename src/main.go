package main

import (
	"fmt"
	"time"
)

type Item struct {
	Id          int
	Title       string
	Price       float64
	CategoryId  int
	Description string
}

type User struct {
	Id       int
	Username string
}

type Category struct {
	Id    int
	Title string
}

type Order struct {
	Id      int
	UserId  int
	ItemId  int
	Date    time.Time
	Address string
}

var intInput int
var Items []Item
var Users []User
var Categories []Category
var Orders []Order

func createUser() {
	var username string
	var id int

	if len(Users) > 0 {
		id = Users[len(Users)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter username:")
	fmt.Scanf("%s\n", &username)

	Users = append(Users, User{
		Id:       id,
		Username: username,
	})
}

func createCategory() {
	var id int
	var category string

	if len(Categories) > 0 {
		id = Categories[len(Categories)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter category:")
	fmt.Scanf("%s\n", &category)

	Categories = append(Categories, Category{
		Id:    id,
		Title: category,
	})
}

func createItem() {
	var id int
	var title string
	var price float64
	var categoryId int
	var description string

	if len(Items) > 0 {
		id = Items[len(Items)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter title:")
	fmt.Scanf("%s,\n", &title)
	fmt.Println("Enter price:")
	fmt.Scanf("%f\n", &price)
	fmt.Println("Enter id of category:")
	fmt.Scanf("%d\n", &categoryId)
	fmt.Println("Enter description:")
	fmt.Scanf("%s\n", &description)

	Items = append(Items, Item{
		Id:          id,
		Title:       title,
		Price:       price,
		CategoryId:  categoryId,
		Description: description,
	})
}

func createOrder() {
	var id int
	var userId int
	var itemId int
	date := time.Now()
	var address string

	if len(Orders) > 0 {
		id = Orders[len(Orders)-1].Id + 1
	} else {
		id = 1
	}

	fmt.Println("Enter id of the user:")
	fmt.Scanf("%d\n", &userId)
	fmt.Println("Enter id of the item:")
	fmt.Scanf("%d\n", &itemId)
	fmt.Println("Enter address:")
	fmt.Scanf("%s\n", &address)

	Orders = append(Orders, Order{
		Id:      id,
		UserId:  userId,
		ItemId:  itemId,
		Date:    date,
		Address: address,
	})
}

func listUsers() {
	for _, user := range Users {
		fmt.Printf("Id: %d\t| Username: %s\n", user.Id, user.Username)
	}
}

func listCategories() {
	for _, category := range Categories {
		fmt.Printf("Id: %d\t| Title: %s\n", category.Id, category.Title)
	}
}

func listItems() {
	for _, item := range Items {
		fmt.Printf("Id: %d\t| Title: %s\t| Price: %f\t| Category Id: %d\t| Description: %s\n",
			item.Id, item.Title, item.Price, item.CategoryId, item.Description)
	}
}

func listOrders() {
	for _, order := range Orders {
		year, month, date := order.Date.Date()
		fmt.Printf("Id: %d\t| User Id: %d\t| Item Id: %d\t| Date: %s\t|, Address: %s\n",
			order.Id, order.UserId, order.ItemId, fmt.Sprintf("%d %s, %d", date, month, year), order.Address)
	}
}

func updateUser() {
	var id int
	var username string

	fmt.Println("Enter user id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter new username:")
	fmt.Scanf("%s\n", &username)

	for i, user := range Users {
		if user.Id == id {
			Users[i].Username = username
			break
		}
	}
}

func updateCategory() {
	var id int
	var categoryTitle string

	fmt.Println("Enter category id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter new category title:")
	fmt.Scanf("%s\n", &categoryTitle)

	for i, category := range Categories {
		if category.Id == id {
			Categories[i].Title = categoryTitle
			break
		}
	}
}

func updateItem() {
	var id int
	var itemTitle string
	var price float64
	var categoryId int
	var description string

	fmt.Println("Enter item id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter new item title:")
	fmt.Scanf("%s\n", &itemTitle)
	fmt.Println("Enter new item price:")
	fmt.Scanf("%f\n", &price)
	fmt.Println("Enter new category id:")
	fmt.Scanf("%d\n", &categoryId)
	fmt.Println("Enter new description:")
	fmt.Scanf("%s\n", &description)

	for i, item := range Items {
		if item.Id == id {
			Items[i].Title = itemTitle
			Items[i].Price = price
			Items[i].CategoryId = categoryId
			Items[i].Description = description
			break
		}
	}
}

func deleteUser() {
	var id int

	fmt.Println("Enter user id:")
	fmt.Scanf("%d\n", id)

	for i, user := range Users {
		if user.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			break
		}
	}
}

func deleteCategory() {
	var id int

	fmt.Println("Enter category id:")
	fmt.Scanf("%d\n", id)

	for i, category := range Categories {
		if category.Id == id {
			Categories = append(Categories[:i], Categories[i+1:]...)
			break
		}
	}
}

func deleteItem() {
	var id int

	fmt.Println("Enter item id:")
	fmt.Scanf("%d\n", id)

	for i, item := range Items {
		if item.Id == id {
			Items = append(Items[:i], Items[i+1:]...)
			break
		}
	}
}

func deleteOrder() {
	var id int

	fmt.Println("Enter order id:")
	fmt.Scanf("%d\n", id)

	for i, order := range Orders {
		if order.Id == id {
			Orders = append(Orders[:i], Orders[i+1:]...)
			break
		}
	}
}

func main() {
	Items = make([]Item, 0)
	Users = make([]User, 0)
	Categories = make([]Category, 0)
	Orders = make([]Order, 0)

LOOP:
	for {
		fmt.Println("What you wanna do? (Select between 0-4).")
		fmt.Println("0. Exit.")
		fmt.Println("1. List data.")
		fmt.Println("2. Create data.")
		fmt.Println("3. Update data.")
		fmt.Println("4. Delete data.")
		fmt.Scanf("%d\n", &intInput)
		switch intInput {

		case 0:
			break LOOP

		case 1:
			fmt.Println("Select the table.")
			fmt.Println("1. Items table.")
			fmt.Println("2. Users table.")
			fmt.Println("3. Categories table.")
			fmt.Println("4. Orders table.")
			fmt.Scanf("%d\n", &intInput)
			switch intInput {
			case 1:
				listItems()

			case 2:
				listUsers()

			case 3:
				listCategories()

			case 4:
				listOrders()

			default:
				fmt.Println("Incorrect input.")
			}

		case 2:
			fmt.Println("Select the table.")
			fmt.Println("1. Items table.")
			fmt.Println("2. Users table.")
			fmt.Println("3. Categories table.")
			fmt.Println("4. Orders table.")
			fmt.Scanf("%d\n", &intInput)
			switch intInput {
			case 1:
				createItem()

			case 2:
				createUser()

			case 3:
				createCategory()

			case 4:
				createOrder()

			default:
				fmt.Println("Incorrect input.")
			}

		case 3:
			fmt.Println("Select the table.")
			fmt.Println("1. Items table.")
			fmt.Println("2. Users table.")
			fmt.Println("3. Categories table.")
			fmt.Scanf("%d\n", &intInput)
			switch intInput {
			case 1:
				updateItem()

			case 2:
				updateUser()

			case 3:
				updateCategory()

			default:
				fmt.Println("Incorrect input.")
			}

		case 4:
			fmt.Println("Select the table.")
			fmt.Println("1. Items table.")
			fmt.Println("2. Users table.")
			fmt.Println("3. Categories table.")
			fmt.Println("4. Orders table.")
			fmt.Scanf("%d\n", &intInput)
			switch intInput {
			case 1:
				deleteItem()

			case 2:
				deleteUser()

			case 3:
				deleteCategory()

			case 4:
				deleteOrder()

			default:
				fmt.Println("Incorrect input.")
			}

		default:
			fmt.Println("Incorrect input. Please, try again.")
		}
	}
}
