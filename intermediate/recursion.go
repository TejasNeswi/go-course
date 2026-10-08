package intermediate

import "fmt"

func main() {

	// fmt.Println(factorial(5))
	// fmt.Println(factorial(10))

	// fmt.Println(sumOfDigits(9))
	// fmt.Println(sumOfDigits(12))
	// fmt.Println(sumOfDigits(12345))

	fmt.Println("0")
	fmt.Println("1")
	fibonacciSeries(10, 3, 1, 0)
}

func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}

	// Recursive case: factorial of n is n*factorial of n-1
	return n * factorial(n-1)

}

func sumOfDigits(n int) int {

	if n < 10 {
		return n
	}

	return n%10 + sumOfDigits(n/10)
}

func fibonacciSeries(n int, i int, prev int, before int) {

	if i > n {
		return
	}
	fmt.Println(prev+before)
	fibonacciSeries(n, i+1, prev+before, prev)
}
