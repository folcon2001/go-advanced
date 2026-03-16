package main

import "fmt"

func main() {
	age := getAge()
	canDrink(age)
	fmt.Println(canDrink(age))
}

func canDrink(age *int) bool {
	return *age >= 18
}

func getAge() *int {
	age := 18
	return &age
}
