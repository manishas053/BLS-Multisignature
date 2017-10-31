package main
import "fmt"
func main() {
  var n int

  fmt.Println("Enter a number")
  fmt.Scanln(&n)
  var a [1000]int

  k:=0
  for i:=2; i<1000; i++ {
    count := 0
    for j:=2; j<=i/2; j++ {
      if i%j == 0{
        count++
      }
    }
    if count==0 {
      a[k] = i
      k++
      //fmt.Println(i)
    }
  }
  fmt.Printf("%dth prime number is :%d\n",n,a[n-1])
}
