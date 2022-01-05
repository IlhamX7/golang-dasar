package main

import "fmt"

func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar go", i)
	// }

	title := "Golang the best language"

	for index, latter := range title {
		fmt.Println("index :", index, "letter", string(latter))
	}
}
