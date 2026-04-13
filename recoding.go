package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read file
	fileName, err := ReadFile()
	if err != nil {
		fmt.Println("error reading file:", err)
	}

	splittedtext := splitByWhiteSpace(fileName)
	converNums := ConvertNums(splittedtext)
	// split by white spaces

	words := ApplyCaseTags(converNums)
	fixedArt := FixArticles(words)
	

}

func FixArticles(words []string) []string {


	for i:= 0; i < len(words); i++{
		lower := strings.ToLower(words[i])
		if lower != "a"{
			continue
		}

		nextword := strings.Trim(words[i+1], ".,!?")

		correct := AorAn(nextword)

		if words[i] == "A"{
			words[i]= capitalize(correct)
		}else{
			words[i]= correct
		}
	}
return words
}

func AorAn(nextWord string) string {
	if nextWord == "" {
		return "a"
	}

	firstLetter := strings.ToLower(string(nextWord[0]))

	switch firstLetter {

	case "a", "e", "i", "o", "u", "h":
		return "an"

	default:
		return "a"
	}
}

func ApplyCaseTags(words []string) []string {
	for index, tag := range words {
		switch tag {
		case "(up)", "(low)", "(cap)":
			if index > 0 {
				words[index-1] = ApplyCase(tag, words[index-1])
			}

		case "(low,", "(up,", "(cap,":
			count, _ := strconv.Atoi(strings.TrimSuffix(words[index+1], ")"))
			start := index - count
			if index < 0 {
				start = 0
			}
			targetWords := words[start:index]

			for i := range targetWords {
				targetWords[i] = ApplyCase(tag, targetWords[i])
			}

		}

	}

	return words

}

func ApplyCase(tag, s string) string {
	switch {
	case strings.HasPrefix(tag, "(low)"):
		return strings.ToLower(s)

	case strings.HasPrefix(tag, "(up)"):
		return strings.ToUpper(s)

	case strings.HasPrefix(tag, "(up)"):
		return Capitalise(s)
	}
	return s
}
func Capitalise(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func ReadFile() (string, error) {
	inputfileName := os.Args[1]
	filecontent, err := os.ReadFile(inputfileName)
	return string(filecontent), err
}

func splitByWhiteSpace(s string) []string {
	return strings.Fields(s)

}

func ConvertNums(words []string) []string {
	for index, tag := range words {
		if index == 0 {
			continue
		}

		switch tag {

		case "(hex)":
			remoevedHex, err := HexToDecimal(words[index-1])
			if err != nil {
				fmt.Println(err)
			}
			words[index-1] = fmt.Sprint(remoevedHex)

		case "(bin)":

			removed, err := BinToDecimal(words[index-1])
			if err != nil {
				fmt.Println(err)
			}
			words[index-1] = fmt.Sprint(removed)
		}

	}

	return words

}

func HexToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func BinToDecimal(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}
