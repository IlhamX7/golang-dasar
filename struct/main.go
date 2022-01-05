package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	IsActive  bool
}

func main() {
	user := User{1, "Zelda", "Safitri", "zelda@nintendo.com", true}
	user2 := User{2, "Link", "Awakening", "link@nintendo.com", true}

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.email)
}
