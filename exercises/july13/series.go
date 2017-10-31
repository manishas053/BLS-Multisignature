//Given a string of digits, output all the contiguous substrings of length `n` in that string.

package main
import "fmt"
func main() {
  var length, subLen int
  fmt.Println("Enter the length of series:")
  fmt.Scanln(&length)
  fmt.Println("Enter the input values:")
  input := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanf("%d", &input[i])
    }
  fmt.Println("Enter the length of substring :")
  fmt.Scanln(&subLen)
  for i := 0; i <= length - subLen; i++ {
    for j := i; j < i + subLen; j++ {
    fmt.Printf("%d ", input[j])
  }
  fmt.Printf("\n")
}
}
