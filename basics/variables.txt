package basics

import "fmt"

var middleName = "Anthony"

func variables() {
	var age int
	var name string = "Mark"
	var name1 = "Bill"

	count := 10
	lastName := "Olsen"

	fmt.Println(age, name, middleName, lastName, count, name1)

	// Default values
	// Numeric types: 0
	// String: ""
	// Boolean: false
	//Pointers, slices, maps, functions, and structs: nil
}

// func printName() {
// 	firstName := "Mark"
// 	fmt.Println(firstName) }
