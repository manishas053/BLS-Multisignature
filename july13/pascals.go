//program to compute pascal's triangle upto a given number of rows

package main
import "fmt"

func main() {
  var rows, coeff int
  fmt.Print("Enter number of rows :")
  fmt.Scanln(&rows)

  for i := 0; i < rows; i++ {
    for space := 1; space <= rows-i; space ++ {
      fmt.Print("  ")
    }
    for j := 0; j <= i; j++ {
      if (i == 0 || j == 0) {
        coeff = 1
      } else {
        coeff = coeff * (i - j + 1) / j
      }
      fmt.Printf("%4d",coeff)
    }
    fmt.Printf("\n")
  }
}
