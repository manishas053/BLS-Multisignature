// Implement a simple shift cipher like Caesar and a more secure substitution cipher

package main

import (
  "fmt"
  "strings"
)

func main() {
var msg string

fmt.Print("Enter Message : ")
fmt.Scanf("%s", &msg)

cipherT := ""
msg = strings.ToUpper(msg)
for i := 0; i < len(msg); i ++ {
  j := ((int(msg[i]) + 3))
  if j > 90 {
    j = j - 26
  }
cipherT = cipherT + string(j)
}

fmt.Println("Cipher text : ", cipherT)

}
