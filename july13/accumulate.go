//Program to Implement the `accumulate` operation, which, given a collection and an operation to
//perform on each element of the collection, returns a new collection containing the result of applying
//that operation to each element of the input collection

package main
import "fmt"

func main() {
  var length int
  fmt.Println("Enter the length:")
  fmt.Scanln(&length)
  fmt.Println("Enter the input values:")
  input := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanln(&input[i])
    }
  fmt.Println(input)

  var op string
  fmt.Printf("Choose any operation :\n1.addition\n2.square\n")
  fmt.Scanln(&op)
  switch op {
  case "addition" :
    for i := 0; i < length; i++ {
      input[i] += input[i]
    }
    fmt.Println(input)

  case "square" :
    for i := 0; i < length; i++ {
      input[i] = input[i] * input[i]
    }
    fmt.Println(input)

  }
}
