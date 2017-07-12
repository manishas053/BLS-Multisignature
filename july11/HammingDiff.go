//Program to find the hamming difference between two DNA strands

package main
import "fmt"
func main() {

    var input,input1 string
    fmt.Print("Enter DNA strand: ")
    fmt.Scanln(&input)
    fmt.Print("Enter another DNA strand: ")
    fmt.Scanln(&input1)

    count := 0
    for i:=0; i<len(input); i++ {
      if input[i] != input1[i]{
        count++
      }
    }
    fmt.Println("Hamming difference between the two DNA strands : ",count)

}
