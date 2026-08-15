package basics

import "fmt"

func main() {

	fruit := "pineapple"

	switch fruit {
	case "apple":
		fmt.Println("It is an apple")
	case "banana":
		fmt.Println("It is a banana")
	default:
		fmt.Println("Unknown fruit")
	}

	day:= "Monday"
	
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Its a week day")
	case "Saturday", "Sunday":
		fmt.Println("Its a week end")
	default:
		fmt.Println("Invalid day")
	}

	number:=15

	switch {
	case number<10:
		fmt.Println("Number is less than 10")
	case number>=10 && number<20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}


	num:=2
	switch {
	case num>1:
		fmt.Println("Greater than 1")
		fallthrough
	case num==2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not 2")
	}

	checkType(10)
	checkType(2.00)
	checkType("hello")
	checkType(true)
}

func checkType (x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("Its an integer")
	case float64:
		fmt.Println("Its float")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Unknown type")
	}
}
