package program

import (
	"encoding/json"
	"fmt"
)

// User description
type User struct {
	ID   int    `json: "userId"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

// AllUsers description
var AllUsers []User

// AddNewUser description
func AddNewUser(u User) (users []User) {
	users = append(users, u)

	return users
}

// DisplayUser description
func DisplayUser(users []User) {
	result, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println(string(result))
	}
}

// Greeting des
func Greeting() string {
	return "Hello world"
}
