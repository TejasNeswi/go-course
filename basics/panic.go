package basics

import "fmt"

func main() {

	// panic(interface{} )

	// Example of a valid input
	// process(10)

	// Example of invalid input
	process(-3)


}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("This is before Panic")
		panic("Input must be a non-negative number")
		// fmt.Println("This is after Panic")

		// fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)

}
