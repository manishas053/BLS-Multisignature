//Program to find acronym of a string

package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	)

func main() {
  var str string
	reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  input, _ := reader.ReadString('\n')
	runes := rune(input[0])
	runes = unicode.ToUpper(runes)
	str = string(runes)
	fmt.Print(str)

	for i := 0; i < len(input); i++ {
    if input[i] == ' ' {
			runes = rune(input[i + 1])
			runes = unicode.ToUpper(runes)
			str = string(runes)
			fmt.Print(str)
    }
  }
  fmt.Printf("\n")
}
