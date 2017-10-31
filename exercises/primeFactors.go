package main
import "fmt"
func main(){
  var num int
  fmt.Println("Enter a number : ")
  fmt.Scanln(&num)
  fmt.Println("Prime factors are :")
  //var count int
  for i:=2; i<=num/2; i++ {
    if num%i == 0 {
      count:= 0
      for j:=2; j<=i/2; j++ {

        if i%j==0{
          count++
        }
      }
      if count == 0 {
        fmt.Println(i)
      }
    }
  }
}
