package main

import "fmt"

func main() {

	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i value", i)
	defer fmt.Println("First Deferred statement executed")
	defer fmt.Println("Second Deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")
	i++
	defer fmt.Println("Deferred i value", i)
	fmt.Println("This is normal execution statement")
	fmt.Println("Value of i", i)
}

