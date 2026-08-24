package basics

import "fmt"

func main() {

	// ... Ellipsis

	// func functionName (param1 type1, param2 type2, param3 ...type3) returnType {
	// 	code block
	// 	return value
	// }

	fmt.Println("Sum of 1, 2, 3: ", sum(1, 2, 3))

	statement, total := sum2("The sum of 1, 2, 3 is", 1, 2, 3)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}
	statement1, total1 := sum2("The sum of 1, 2, 3, 4, 5, 9 is ", numbers...)
	fmt.Println(statement1, total1)

}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func sum2(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
