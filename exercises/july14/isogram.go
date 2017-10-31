//Program to determine if a word is an isogram

package main
import (
  "fmt"
  "bufio"
  "os"
  )
func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter a word : ")
  word, _ := reader.ReadString('\n')
  count := 0
  for i := 0; i < len(word); i++ {
    for j := 0; j < len(word); j++ {
      if i != j {
        if word[i] == word[j] {
          count ++
        }
      }
    }
  }
  if count == 0 {
    fmt.Println("Isogram")
  } else {
    fmt.Println("Not an Isogram")
  }
}
