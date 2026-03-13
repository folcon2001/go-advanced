package main

import "fmt"
type User struct {
	Name string
}
func main() {
user := &User{
	Name: "Dfcz",
}

	fmt.Println(user)
}
