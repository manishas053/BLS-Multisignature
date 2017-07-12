//Program to find palindrome products in a given range

package main
import "fmt"
func main() {
  var a,b,product,num,r int
  fmt.Println("Enter the range :")
  fmt.Scanln(&a,&b)
  fmt.Println(a,b)
  for i:=a; i<=b; i++ {
    for j:=a; j<=b; j++ {
      product = i*j
      num = product
      sum:= 0
      for k:=num; k>0; k=k/10 {
        r = k%10
        sum = sum*10 + r
      }
      if sum== product{
        fmt.Printf("%d * %d = %d\n",i,j,product)
      }
    }
  }
}
