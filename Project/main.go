package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}

var users []User

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add User")
		fmt.Println("2. Fetch Users")
		fmt.Println("3. Fetch User by ID")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		choiceStr, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceStr)

		switch choice {
		case "1":
			addUser(reader)
		case "2":
			fetchUsers("")
		case "3":
			fetchUserByID(reader)
		case "4":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addUser(reader *bufio.Reader) {
	fmt.Println("\nAdding a new user:")
	fmt.Print("Enter User name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter user email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter user phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Enter user address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	newUser := User{
		ID:      len(users) + 1,
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
	}

	users = append(users, newUser)
	fmt.Println("User added successfully!")
}

func fetchUsers(userID string) {
	if len(users) == 0 {
		fmt.Println("No users added yet.")
		return
	}

	fmt.Println("\nList of Users:")
	for _, user := range users {
		if userID != "" && strconv.Itoa(user.ID) != userID {
			continue
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s, Address: %s\n", user.ID, user.Name, user.Email, user.Phone, user.Address)
	}
}

func fetchUserByID(reader *bufio.Reader) {
	fmt.Print("Enter user ID: ")
	userIDStr, _ := reader.ReadString('\n')
	userIDStr = strings.TrimSpace(userIDStr)

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("Invalid user ID. Please enter a valid number.")
		return
	}

	found := false
	for _, user := range users {
		if user.ID == userID {
			found = true
			fetchUsers(userIDStr)
			break
		}
	}
	if !found {
		fmt.Println("User not found.")
	}
}
