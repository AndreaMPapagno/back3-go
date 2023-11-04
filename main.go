package main

import "fmt"

type User struct {
	Name     string
	Lastname string
	Age      int
	Email    string
	Password string
}

func (u *User) changeName(newName string) {
	u.Name = newName
}

func (u *User) changeAge(newAge int) {
	u.Age = newAge
}

func (u *User) changeEmail(newEmail string) {
	u.Email = newEmail
}

func (u *User) changePass(newPass string) {
	u.Password = newPass
}

func main() {
	user1 := User{
		Name:     "Gael",
		Lastname: "Charles",
		Age:      11,
		Email:    "galuCha@mail.com",
		Password: "abc123",
	}
	fmt.Printf("El nombre del usuario es %s, su edad es %d, mail es %s y su pass es %s\n", user1.Name, user1.Age, user1.Email, user1.Password)
	user1.changeName("Lara")
	user1.changeAge(18)
	user1.changeEmail("laruch@mail.com")
	user1.changePass("def456")
	fmt.Printf("El nombre del usuario es %s, su edad es %d, mail es %s y su pass es %s", user1.Name, user1.Age, user1.Email, user1.Password)
}
