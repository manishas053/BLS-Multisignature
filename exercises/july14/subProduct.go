//Given a string of digits, calculate the largest product for a contiguous substring of digits of length n


package main
import "fmt"
func main() {
  var length, subLen, product int
  var pArray []int
  fmt.Println("Enter the length of series:")
  fmt.Scanln(&length)
  fmt.Println("Enter the input values:")
  input := make([]int, length)
  for i := 0; i < length; i++ {
      fmt.Scanf("%d", &input[i])
  }
  fmt.Println("Enter the length of substring :")
  fmt.Scanln(&subLen)
  fmt.Printf("The substings are : ")
  for i := 0; i <= length - subLen; i++ {
    product = 1
    for j := i; j < i + subLen; j++ {
    fmt.Printf("%d", input[j])
    product = product * input[j]

  }
pArray = append(pArray, product)
  fmt.Printf(" ")

}
fmt.Printf("\nThe products of substrings : ")
fmt.Println(pArray)
max := pArray[0]
for i := 0; i < subLen; i++ {
  if pArray[i] < pArray[i + 1] {
    max = pArray[i + 1]
  }
}
fmt.Printf("Largest product : ")
fmt.Println(max)
}
