package main
import "fmt"
func main() {
  var num int
  fmt.Println("Enter a number :")
  fmt.Scanln(&num)
  var sum int
  for i:=1; i<num; i++ {
    if num%i == 0 {
      sum = sum + i
    }
  }
  if sum==num{
    fmt.Println("Perfect number")
  }
  if sum < num {
    fmt.Println("Deficient")
  }
  if sum > num {
    fmt.Println("Abundant")
  }
}
