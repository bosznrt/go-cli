package main

import (
	"log"
	"os"
	"program"
	"strconv"
	"strings"
)

func main() {
	operation := os.Args[1]

	var name string
	var age string
	var iAge int

	if len(os.Args) > 2 && operation == "add" {
		tempName := strings.Split(os.Args[2], "=")
		if tempName[0] != "name" {
			log.Fatalf("Error : Invalid argument unexpected name")
		}
		name = tempName[1]
		tempAge := strings.Split(os.Args[3], "=")
		if tempAge[0] != "age" {
			log.Fatalf("Error : Invalid argument unexpected age")
		}
		age = tempAge[1]

		iAge, _ = strconv.Atoi(age)
	}

	AllUsers := program.ReadStore()

	switch operation {
	case "add":
		size := len(AllUsers)
		currentID := 1

		if size > 0 {
			currentID = AllUsers[len(AllUsers)-1].ID + 1
		}

		AllUsers = program.AddNewUser(program.User{ID: currentID, Name: name, Age: iAge})

		program.WriteStore(AllUsers)

	case "list":
		program.DisplayUser(AllUsers)
	}
}
