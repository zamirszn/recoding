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
	fmt.Println(fixPunctuation("Fix the punctuations , in this code ! , notice the space between the punctuations"))
}

func hexToDecimal(hexStr string) (int64, error) {
	con, err := strconv.ParseInt(hexStr, 16, 64)
	return con, err
}

func binToDecimal(bin string) (int64, error) {
	con, err := strconv.ParseInt(bin, 2, 64)
	return con, err
}

func capitalize(word string) string {
	cap := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	return cap
}

func capitalizeTheNth(words []string, n int) []string {
	start := len(words) - n
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

func fixPunctuation(word string) string {
	word = strings.ReplaceAll(word, " ,", ",")
	word = strings.ReplaceAll(word, " !", "!")
	return strings.TrimSpace(word)
}
