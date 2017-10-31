//Given a word and a list of possible anagrams, select the correct sublist

package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  )

func main() {
  var word, index string
  var not_found int

  fmt.Print("Enter a word : ")
  fmt.Scanln(&word)

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("List possible angrams : ")
  input, _ := reader.ReadString('\n')

  count := 0
  for i := 0; i < len(input); i ++ {
    if input[i] == ' ' {
      count ++
    }
  }

  fmt.Println("Anagram : ")
  s := strings.Split(input, " ")
  for j :=0; j < count+1; j ++ {
    index = s[j]

    if len(word) == len(s[j]) {
      len := len(word)
      for i := 0; i <= len; i ++ {
        found := 0
        for k := 0; k <= len; k ++ {
          if (word[j] == index[k]) {
            found = 1
            break
          }
        }
        if found == 0 {
          not_found = 1
          break
        }
      }
      if not_found == 0 {
        fmt.Println("    ",index)
      }
    }
  }
}
