package basics

import "fmt"

// global variables (accessible to all functions)
// middleName := "Cane"  // := can only be used inside the functions
var middleName = "Cane"

func main() {
	var age int
	var name1 string = "John"
	var name2 = "Jane"

	count := 10
	lastName := "Smith"

	fmt.Println(middleName)

	// Default values
	// Numeric types: 0
	// Boolean types: false
	// String types: ""
	// Pointers, slices, maps, functions and structs: nil

	// --- SCOPE

	fmt.Println(firstName)

}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName)
}