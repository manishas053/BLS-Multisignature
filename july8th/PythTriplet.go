package main
import "fmt"
func main() {

  for a:=1 ; a<=500; a++{            //loop for base
    for b:=1; b<=500; b++{           //loop for altitude
      for c:=1; c<=500; c++ {        //loop for hypotenuse
        if ((a*a)+(b*b)==(c*c)) {    //pythagorus theorem
          if (a < b) {               // base < altitude
          if a + b+ c == 1000{ 
            fmt.Printf("Product of %d %d %d is : %d\n",a,b,c, a*b*c)
          }
        }
        }
      }
    }
  }
}
