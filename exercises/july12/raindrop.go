//Program to Convert a number to a string, the contents of which depend on the number's factors

package main
import "fmt"
func main() {
  var num int
  var factor [10]int
  j:=0
  count:=0
  fmt.Println("Enter a number")
  fmt.Scanln(&num)
  for i:=1; i<=num; i++ {
    if num%i == 0 {
      factor[j] = i
      j++
      if i == 3 {
        fmt.Printf("Pling")
      } else if i == 5 {
        fmt.Printf("Plang")
      } else if i == 7 {
        fmt.Printf("Plong")
      }
    }
  }
  for j:=0; j<10; j++ {
    if (factor[j]==3 || factor[j]==5 || factor[j]==7) {
      count++
      }
  }
  if count==0 {
    fmt.Println(num)
  }
  fmt.Printf("\n")
}
