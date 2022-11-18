package main

import (
	"fmt"
)

type UserData struct {
	ID       int
	Name     string
	Location string
}

func (user *UserData) GreetUser() string {
	return fmt.Sprintf("Hello %s from %s",
		user.Name, user.Location)
}

func CreateNewUser(id int, name, location string) *UserData {
	id++
	return &UserData{id, name, location}
}

func main() {
	user := CreateNewUser(10, "Maryam", "PK")
	fmt.Println(user.GreetUser())
}
