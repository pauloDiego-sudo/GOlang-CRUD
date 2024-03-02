### Code Explanation

#### Package Declaration
```go
package main
```
Defines the package name. Here, `main` indicates that this file is a standalone executable program.

#### Import Statements
```go
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
```
Imports necessary packages for input/output operations, string manipulation, and converting strings to integers.

#### Database File Constant
```go
const dbFile = "users.txt"
```
Specifies the file name where user data is stored.

#### User Struct Definition
```go
type User struct {
	ID       string
	Name     string
	Function string
	Age      int
}
```
Defines a `User` type with ID, Name, Function, and Age as attributes.

#### CreateUser Function
```go
func CreateUser(name string, function string, age int) {
```
Defines a function to create a new user and append their data to the file.

```go
	file, _ := os.OpenFile(dbFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
```
Opens `users.txt` for appending, creating the file if it doesn't exist. Ensures the file is closed after the function executes.

```go
	newID := generateUniqueID()
	newUser := User{ID: newID, Name: name, Function: function, Age: age}
```
Generates a unique ID for the new user and creates a `User` instance with the provided data.

```go
	line := fmt.Sprintf("%s,%s,%s,%d\n", newUser.ID, newUser.Name, newUser.Function, newUser.Age)
	_, _ = file.WriteString(line)
	fmt.Println("User created:", newUser)
}
```
Formats the user's data as a string and writes it to the file. Prints a confirmation message.

#### ReadUsers Function
```go
func ReadUsers(id string) []User {
```
Defines a function to read users from the file, either a specific user by ID or all users if no ID is provided.

```go
	file, _ := os.Open(dbFile)
	defer file.Close()
```
Opens `users.txt` for reading and ensures the file is closed afterwards.

```go
	var users []User
	scanner := bufio.NewScanner(file)
```
Initializes a slice to store users and a scanner to read the file line by line.

```go
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		idade, _ := strconv.Atoi(parts[3])
		user := User{ID: parts[0], Name: parts[1], Function: parts[2], Age: idade}
		if id == "" || user.ID == id {
			users = append(users, user)
		}
	}
	return users
}
```
Loops through each line in the file, splits the line into parts, converts the age to an integer, creates a `User` instance, and adds it to the slice if it matches the ID criteria.

#### DisplayUsers Function
```go
func DisplayUsers(users []User) {
```
Defines a function to print user data in a formatted table.

```go
  fmt.Println("Displaying User(s):")
  fmt.Printf("%-2s %-15s %-10s %s\n", "ID", "Name", "Function", "Age")
```
Prints the table header.

```go
  for _, user := range users {
    fmt.Printf("%-2s %-15s %-10s %d\n", user.ID, user.Name, user.Function, user.Age)
  }
}
```
Iterates over the slice of users and prints each user's data.

#### UpdateUser Function
Defines a function to update a user's data by ID.

#### DeleteUser Function
Defines a function to delete a user by ID.

#### generateUniqueID Function
```go
func generateUniqueID() string {
```
Generates a unique ID for a new user based on the current number of users.

#### main Function
```go
func main() {
```
Contains the main program logic, displaying a menu and executing actions based on user input.

#### showMenu Function
```go
func showMenu() {
```
Displays the available actions to the user.

#### getUserInput Function
```go
func getUserInput(prompt string) string {
```
Prompts the user for input and returns the input as a string.

