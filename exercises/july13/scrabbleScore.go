//Program to compute the scrabble score for a given word

package main
import "fmt"
func main() {
  var word string
  fmt.Print("Enter a word : ")
  fmt.Scanln(&word)
  count := 0
  for i := 0; i < len(word); i++ {
    if ((word[i] == 'a') || (word[i] == 'e') || (word[i] == 'i') || (word[i] == 'o') || (word[i] == 'u') || (word[i] == 'l') || (word[i] == 'n') || (word[i] == 'r') || (word[i] == 's') || (word[i] == 't')) {
      count = count + 1
      } else if ((word[i] == 'd') || (word[i] == 'g')) {
        count = count + 2
      } else if ((word[i] == 'b') || (word[i] == 'c') || (word[i] == 'm') || (word[i] == 'p')) {
        count = count + 3
      } else if ((word[i] == 'f') || (word[i] == 'h') || (word[i] == 'v') || (word[i] == 'w') || (word[i] == 'y')) {
        count = count + 4
      } else if ((word[i]) == 'k') {
        count = count + 5
      } else if ((word[i] == 'j') || (word[i] == 'x')) {
        count = count + 8
      } else if ((word[i] == 'q') || (word[i] == 'z')) {
        count = count + 10
      }
  }
  fmt.Println("Scrabble score : ",count)
}
