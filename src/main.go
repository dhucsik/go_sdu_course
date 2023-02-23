package main

import (
	"flag"
	"fmt"
	"module/config"
	"module/pkg/db"
	"module/pkg/services"
	"os"
)

var intInput int
var strInput string

func main() {
	var username string
	var password string

	flag.StringVar(&username, "username", "postgres", "Name of the user")
	flag.StringVar(&password, "password", "", "Password")
	flag.Parse()

	database, err := db.Connect(username, password)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	config.App.DB = database
	defer config.App.DB.Close()

LOOP1:
	for {
		fmt.Println("Do you have an account?(Yes or No)")
		fmt.Scanf("%s\n", &strInput)

		switch strInput {

		case "Yes":
			login()
			break LOOP1

		case "No":
			register()
			break LOOP1

		default:
			fmt.Println("Invalid input. Try again.")
		}
	}

LOOP2:
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
			break LOOP2

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
				fmt.Println(config.App.User)

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

func login() {
	for {
		fmt.Println("Enter your username:")
		fmt.Scanf("%s\n", &strInput)

		var password string

	OuterLoop:
		for _, user := range db.Users {
			if user.Username == strInput {
				for i := 0; i < 3; i++ {
					fmt.Println("Enter your password:")
					fmt.Scanf("%s\n", &password)
					if password == user.Password {
						config.App.User = user
						break OuterLoop
					} else {
						fmt.Println("Incorrect password. Try again.")
					}
				}
				fmt.Println("You spent all 3 attempts. Try again.")
				os.Exit(0)
			}
		}

		fmt.Println("User with such a username does not exist. Try again.")
		fmt.Println("If you want to exit, enter the number 0. Otherwise type anything.")
		fmt.Scanf("%d\n", &intInput)

		if intInput == 0 {
			os.Exit(0)
		}

	}
}

func register() {
	config.App.User = services.CreateUser()
}
