package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg: CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// Eg: user_id, first_name, http_request
	// variable, file names

	// UPPERCASE
	// constants

	// mixedCase
	// Eg: javaScript, htmlDocument, isValid

	const MAXRETRIES = 5
	var employeeId = 1001
	fmt.Println("EmployeeID: ", employeeId)

}
