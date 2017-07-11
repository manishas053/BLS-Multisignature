package main
import "fmt"

func main() {
  var limit int
  var num1 int
  var num2 int
  fmt.Println("Enter the limit")
  fmt.Scanln(&limit)
  fmt.Println("Enter the 1st number :")
  fmt.Scanln(&num1)
  fmt.Println("Enter the 2nd number :")
  fmt.Scanln(&num2)

  sum := 0
  for i:=1; i<limit; i++ {
    if (i%num1==0 || i%num2==0) {
      sum = sum + i
    }
  }
  fmt.Println("Sum of multiples is : ", sum)
}
