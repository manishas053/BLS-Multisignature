package main

import "fmt"

func main() {
  var limit int
  fmt.Println("Enter the limit :")
  fmt.Scanln(&limit)

  //var sumSq int
  //var sum int
  var Sqsum int
  var diff int

  sumSq :=0
  sum := 0
  //var i int
  for i:= 1; i<=limit; i++ {
    sumSq = sumSq + (i*i)
    sum = sum + i
  }
  Sqsum = sum * sum
  diff = Sqsum - sumSq
  fmt.Println("Difference is : ",diff)
}
