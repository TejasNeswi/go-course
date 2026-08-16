package main

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariable map[keyType]valueType

	// Using make
	// mapVariable := make(map[keyType]valueType)

	// Using map literal
	// mapVariable = map[keyType]valueType {
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"]) // In case of a non existant key we get the default value of the valueType i.e 0

	myMap["code"] = 21
	fmt.Println(myMap)

	// Delete a key
	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, unknownValue := myMap["key1"]
	// fmt.Println(value)
	fmt.Println("Is a value associated with key1:", unknownValue)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)
	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	var myMap4 map[string]string

	fmt.Println(myMap4)

	if myMap4==nil {
		fmt.Println("The map is initialized to nil")
	} else {
		fmt.Println("The map is not initialized to nil")
	}

	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "value"
	// fmt.Println(myMap4)

	myMap4= make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	fmt.Println("The lenght of myMap is ", len(myMap))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
