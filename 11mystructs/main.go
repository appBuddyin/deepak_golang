package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	deepak := User{"deepak", "deepak@appbuddy.in", true, 16}
	fmt.Println(deepak)
	fmt.Printf("hitesh details are: %+v\n", deepak)
	fmt.Printf("Name is %v and email is %v.", deepak.Name, deepak.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
