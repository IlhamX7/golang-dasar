package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["JavaScript"] = 8
	// myMap["go"] = 10

	// fmt.Println(myMap["ruby"])

	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"JavaScript": "hype",
	}

	// for key, value := range myMap {
	// 	fmt.Println("key : ", key, " value: ", value)
	// }

	// fmt.Println("==========")

	// delete(myMap, "ruby")

	// for key, value := range myMap {
	// 	fmt.Println("Key : ", key, " value: ", value)
	// }

	value, isAvailable := myMap["go"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}
