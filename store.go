package program

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ReadStore description
func ReadStore() (users []User) {
	file, errRead := ioutil.ReadFile("result.json")
	if errRead != nil {
		_ = ioutil.WriteFile("result.json", []byte(""), 0644)
	}

	_ = json.Unmarshal(file, &users)

	return users
}

// WriteStore description
func WriteStore(users []User) []byte {

	result, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		_ = ioutil.WriteFile("result.json", result, 0644)
	}

	return result
}
