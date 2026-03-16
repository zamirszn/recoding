package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(binToDecimal("10"))
	fmt.Println(capitalize("hELLo"))
	fmt.Println(capitalizeTheNth([]string{"go", "is", "really", "fun"}, 3))
	fmt.Println(joinWithPunctuation([]string{"hello", ",", "world", "!"}))
	fmt.Println(isPunctuation("!"))
	fmt.Println(isPunctuation("x"))
	fmt.Println(aOrAn("apple"))
	fmt.Println(fixArticles("There it was. A amazing rock. A honest man. A book."))
	// 
	result := fixSingleQuotes("' hello world '")
	fmt.Printf("%q\n", result)
	
}

// 1
func hexToDecimal(hexStr string) (int64, error) {
	con, err := strconv.ParseInt(hexStr, 16, 64)
	return con, err
}

// 2
func binToDecimal(bin string) (int64, error) {
	con, err := strconv.ParseInt(bin, 2, 64)
	return con, err
}

// 3
func capitalize(word string) string {
	cap := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	return cap
}

// 4
func capitalizeTheNth(words []string, n int) []string {
	start := len(words) - n
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

// 5
func joinWithPunctuation(tokens []string) string {
    result := strings.Join(tokens, "")

    // Replace the comma with a comma and a space
    return strings.ReplaceAll(result, ",", ", ")
}

// 6
func isPunctuation(s string) bool {
	return s == "," || s == "!"
}

// 7
func aOrAn(nextWord string) string {
	if nextWord == "" {
		return "a"
	}

	first := strings.ToLower(string(nextWord[0]))

	switch first {
	case "a", "e", "i", "o", "u", "h":
		return "an"
	default:
		return "a"
	}
}

// 8
func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "A" || words[i] == "a" {

			next := strings.Trim(words[i+1], ".,!?")

			article := aOrAn(next)

			if words[i] == "A" {
				words[i] = strings.ToUpper(article[:1]) + article[1:]
			} else {
				words[i] = article
			}
		}
	}

	return strings.Join(words, " ")
}

// 9
func fixSingleQuotes(text string) string {
	text = strings.Trim(text, "'")
	text = strings.TrimSpace(text)
	return "'" + text + "'"
}

