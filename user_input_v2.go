package main

import (
	"fmt"
	"strings"
)

func wordProcess() {
	var word string
	var operation string
	var index int
	for {
		fmt.Print("\nEnter a word: ")
		fmt.Scan(&word)
		word = strings.TrimSpace(word)
		if word == "" {
			fmt.Println("Word cannot be empty!")
			continue
		}
		for {
			fmt.Println("\n[1] lastchar  [2] capitalism  [3] delete index  [4] quite")
			fmt.Print("Choose operator: ")
			fmt.Scanln(&operation)
			operation = strings.TrimSpace(strings.ToUpper(operation))

			switch operation {
			case "1", "LASTCHAR":
				fmt.Printf("last character: '%c'\n", word[len(word)-1])

			case "2", "CAPITALISM":
				fmt.Printf("capitalize: '%s'\n", strings.ToUpper(word))

			case "3", "DELETE INDEX":
				fmt.Printf("Enter index (0-%d): ", len(word)-1)
				_, err := fmt.Scanln(&index)
				if err != nil || index < 0 || index >= len(word) {
					fmt.Println("Error try again")
					break
				}

				fmt.Printf("Result: '%s'\n", word[:index]+word[index+1:])

			case "4", "QUITE":
				fmt.Println("GOODBYE!")
				return

			default:
				fmt.Println("Invalid choice")
				continue
			}
			break
		}
	}
}

func main() {
	wordProcess()
}
