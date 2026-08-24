package basics

import "fmt"

func main() {

	// func <name> (parameters list) returnType {
	// 	code Block
	// 	return value
	// }

	// sum := add(1, 2)

	fmt.Println(add(2, 3))

	greet := func() {
		fmt.Println("hello anonymous function")
	}
	greet()

	operation := add
	result := operation(3, 5)
	fmt.Println(result)

	// Passing a func as an argument
	result1 := applyOperation(5, 3, add)
	fmt.Println("5 + 3 = ",result1)

	// Returning and using a func
	multiplyBy2:=createMultiplier(2)
	fmt.Println("6 x 2 = ", multiplyBy2(6))
}
func add(a, b int) int {
	return a + b
}

// This is a func that takes a func as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Func that returns a func
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
