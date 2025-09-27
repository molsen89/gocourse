package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//  PascalCase (PascalCase for struct fields and types)
	// e.g., MyFunction, MyVariable
	//  camelCase (variables, functions)
	// e.g., myFunction, myVariable
	//  snake_case (file_names, variables)
	// e.g., my_function, my_variable
	//  kebab-case (not used in Go)
	// e.g., my-function, my-variable
	//  UPPER_SNAKE_CASE (for constants)
	// e.g., MY_CONSTANT, MAX_VALUE

	const MAXRETRIES = 3
	var EmployeeID = 1001
	fmt.Println("Employee ID:", EmployeeID)

}
