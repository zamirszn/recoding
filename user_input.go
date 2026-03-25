package main

import (
	"fmt"
	"strings"
)

func main() {
	var userinput string
	var operation string
	var userindex int

	fmt.Println("input a word")
	fmt.Scanln(&userinput)

	fmt.Println("input an operation")
	fmt.Scanln(&operation)
	switch operation {
	case "lastword":
		fmt.Println(lastLetter(userinput))
	case "capitalise":
		fmt.Println(capitalize(userinput))
	case "delete":
		fmt.Println("index to delete")
		fmt.Scanln(&userindex)
		delete(userinput, userindex)
	default:
		fmt.Println("invalid operation")
	}
}

func lastLetter(s string) string {
	l := []rune(s)
	return string(l[len(l)-1])
}

func capitalize(s string) string {
	return strings.ToUpper(s)
}

func delete(s string, index int) {
	lenofword := len(s)

	if index > lenofword || index < 1 {
		fmt.Printf("Your index in out of range %v, range should be within 1 and %v \n", index, lenofword)
		// restart code
		main()
	}
	empty := []rune{}
	for i, char := range s {
		if i != index {
			empty = append(empty, char)
		}
	}

	fmt.Println(string(empty))
}
