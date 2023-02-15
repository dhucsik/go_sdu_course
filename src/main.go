package main

import (
	"fmt"
	"module/services"
)

var intInput int

func main() {
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
				services.ListItems()

			case 2:
				services.ListUsers()

			case 3:
				services.ListCategories()

			case 4:
				services.ListOrders()

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
				services.CreateItem()

			case 2:
				services.CreateUser()

			case 3:
				services.CreateCategory()

			case 4:
				services.CreateOrder()

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
				services.UpdateItem()

			case 2:
				services.UpdateUser()

			case 3:
				services.UpdateCategory()

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
				services.DeleteItem()

			case 2:
				services.DeleteUser()

			case 3:
				services.DeleteCategory()

			case 4:
				services.DeleteOrder()

			default:
				fmt.Println("Incorrect input.")
			}

		default:
			fmt.Println("Incorrect input. Please, try again.")
		}
	}
}
