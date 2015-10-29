package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
}

func main() {
	user1 := User{Name: "Frank"}
	user2 := User{Name: "Bob"}

	users := []User{}
	users = append(users, user1)
	users = append(users, user2)
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
