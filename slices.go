package main

import (
	"fmt"
	"slices"
)

func main() {

	// var sliceName[] ElementType

	// var numbers []int

	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5) // we are initializing a slice of capacity 5

	a := [5]int{1, 2, 3, 4, 5}

	slice1 := a[1:4] // index 4 is not included, it will store from index 1 to index 3. This is extracting a slice from an array

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice1", slice1)

	sliceCopy := make([]int, len(slice1))

	copy(sliceCopy, slice1)
	fmt.Println("SliceCopy", sliceCopy)

	// var nilSlice []int // its a nil slice

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice1", slice1[3])

	// slice1[3]=50
	// fmt.Println("Element at index 3 of slice1", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	}

	// Multi-Dimensional slice

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i]=make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j]=i+j
		}
	}

	fmt.Println(twoD)

	// slice[low:high]

	slice2:= slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The length of slice2 is", len(slice2))
}
