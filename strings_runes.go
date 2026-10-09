package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!" // modifying in the same line. starts from start
	message3 := "Hello, Go!"   // modifying in the same line. starts from start
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message var is", len(message))
	fmt.Println("Length of message3 var is", len(message3))
	fmt.Println("Length of rawMessage var is", len(rawMessage))

	fmt.Println("The first char in message var is", message[0]) // ASCII value

	greeting := "Hello"
	name := "Alice"

	msg := greeting + name

	fmt.Println(msg)

	str1 := "Apple"
	str2 := "Banana"
	str3:= "app" // a has an ASCII value of 97 and A has ASCII of 65
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c \n", i, char)
		fmt.Printf("%x\n", char)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))
	fmt.Println("Length of string: ",len(greeting) )


	greetingWithName:= greeting+name

	fmt.Println(greetingWithName)

	// RUNES

	var ch rune = 'a'

	fmt.Println(ch)
	fmt.Printf("%c\n", ch)

	cstr:= string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T", cstr)

}
