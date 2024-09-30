package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var emoji string
	fmt.Scanln(&emoji)

	// Call the `utf8.RuneCountInString()` function below to count
	// how many runes are in the `emoji` string variable:
	runeCount := utf8.RuneCountInString(emoji)

	fmt.Println(runeCount) // DO NOT delete this line; it prints the `emoji` rune count
}
