package main

import "fmt"

func main() {
	languages := [...]string{
		"Ruby",
		"Python",
		"JavaScript",
		"Go",
		"C#",
	}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	for index, lang := range languages {
		fmt.Println("index", index, "languages : ", lang)
	}

}
