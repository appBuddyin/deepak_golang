package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	deepak := User{"deepak", "deepak@appbuddy.in", true, 16}
	fmt.Println(deepak)
	fmt.Printf("deepak details are: %+v\n", deepak)
	fmt.Printf("Name is %v and email is %v.\n", deepak.Name, deepak.Email)
	deepak.GetStatus()
	deepak.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", deepak.Name,deepak.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@appbuddy.in"
	fmt.Println("Email of this user is: ", u.Email)
}
