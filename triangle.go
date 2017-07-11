package main

import "fmt"

func main() {
  fmt.Println("Enter the first side of triangle :")
  var a int
  fmt.Scanln(&a)
  //fmt.Println(a)
  fmt.Println("Enter the second side of triangle :")
  var b int
  fmt.Scanln(&b)
  fmt.Println("Enter the third side of triangle :")
  var c int
  fmt.Scanln(&c)

  if a==b && b==c {
    fmt.Println("Equilateral triangle")
  }else if (a==b || b==c || (a==b && b==c)) {
    fmt.Println("Isosceles triangle")
  }else if (a!=b && b!=c) {
    fmt.Println("Scalene triangle")
  }
}
