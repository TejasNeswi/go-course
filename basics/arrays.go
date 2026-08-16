package basics

import "fmt"

func main() {

	// var arrayName[size] elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}

	fmt.Println("Fruits array: ", fruits)

	fmt.Println("Third Element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray
	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index", i, "is", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// underscore is blank identifier, used to store unused values
	for _, v := range numbers {
		fmt.Println("Value: ", v)
	}

	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	// c := 3
	// _ = c

	fmt.Println("The length of the array numbers is", len(numbers))


	// Comparing Arrays
	array1:=[3]int{1,2,3}
	array2:=[3]int{10,2,3}
	fmt.Println("Array1 is equal to Array 2:",array1==array2)


	var matrix [3][3] int=[3][3]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println(matrix)

	// If we want to point the copiedArray to the originalArray
	originalArrayP := [3]int{1, 2, 3}
	var copiedArrayP * [3]int // Stores the Address of an array that can store 3 integers
	copiedArrayP = &originalArrayP

	copiedArrayP[0] = 100

	fmt.Println("Original Array:", originalArrayP)
	fmt.Println("Copied Array:", *copiedArrayP)
}

func someFunction() (int, int) {
	return 1, 2
}
