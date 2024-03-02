package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dbFile = "users.txt"

// User represents a user in the system
type User struct {
	ID       string
	Name     string
	Function string
	Age      int
}

// Create a new user and append it to the file
func CreateUser(name string, function string, age int) {
	file, _ := os.OpenFile(dbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	newID := generateUniqueID()
	newUser := User{ID: newID, Name: name, Function: function, Age: age}

	line := fmt.Sprintf("%s,%s,%s,%d\n", newUser.ID, newUser.Name, newUser.Function, newUser.Age)
	_, _ = file.WriteString(line)
	fmt.Println("User created:", newUser)
}

// Read all users or a single user by ID from the file
func ReadUsers(id string) []User {
	file, _ := os.Open(dbFile)
	defer file.Close()

	var users []User
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",") //Tokenize the line
		idade, _ := strconv.Atoi(parts[3])
		user := User{ID: parts[0], Name: parts[1], Function: parts[2], Age: idade}
		if id == "" || user.ID == id {
			users = append(users, user)
		}
	}
	return users
}

func DisplayUsers(users []User) {
  fmt.Println("Displaying User(s):")
  fmt.Printf("%-2s %-15s %-10s %s\n", "ID", "Name", "Function", "Age")
  for _, user := range users {
    fmt.Printf("%-2s %-15s %-10s %d\n", user.ID, user.Name, user.Function, user.Age)
  }
}


// Update a user's name by ID
func UpdateUser(id, newName string, newFunc string, newAge int) {
	users := ReadUsers("")
	file, _ := os.Create(dbFile) // This will overwrite the existing file
	defer file.Close()

	for _, user := range users {
		if user.ID == id {
			user.Name = newName
      user.Function = newFunc
      user.Age = newAge
		}
		line := fmt.Sprintf("%s,%s,%s,%d\n", user.ID, user.Name,user.Function,user.Age)
		_, _ = file.WriteString(line)
	}
}

// Delete a user by ID
func DeleteUser(id string) {
	users := ReadUsers("")
	file, _ := os.Create(dbFile) // This will overwrite the existing file
	defer file.Close()

	for _, user := range users {
		if user.ID != id {
      line := fmt.Sprintf("%s,%s,%s,%d\n", user.ID, user.Name,user.Function,user.Age)
			_, _ = file.WriteString(line)
		}
	}
}

// A very simple function to generate unique IDs.
func generateUniqueID() string {
	return fmt.Sprintf("%d", len(ReadUsers(""))+1)
}

func main() {
	for {
		showMenu()
		action := getUserInput("Choose an action: ")

		switch action {
		case "1":
			name := getUserInput("Enter user name: ")
      function := getUserInput("Enter user function: ")
      age := getUserInput("Enter user age: ")
      idade,_ := strconv.Atoi(age)
			CreateUser(name,function,idade)
		case "2":
			id := getUserInput("Enter user ID to find: ")
			users := ReadUsers(id)
      DisplayUsers(users)
		case "3":
			allUsers := ReadUsers("")
      DisplayUsers(allUsers)
		case "4":
			id := getUserInput("Enter user ID to update: ")
			newName := getUserInput("Enter new name: ")
      newFunction := getUserInput("Enter new function: ")
      newAge := getUserInput("Enter new age: ")
      newIdade,_ := strconv.Atoi(newAge)      
			UpdateUser(id, newName,newFunction,newIdade)
		case "5":
			id := getUserInput("Enter user ID to delete: ")
			DeleteUser(id)
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\nUser Management System")
	fmt.Println("1) Create User")
	fmt.Println("2) Find User")
	fmt.Println("3) Display Users")
	fmt.Println("4) Update User")
	fmt.Println("5) Delete User")
	fmt.Println("6) Exit")
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
