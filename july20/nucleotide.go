//Given a DNA string, compute how many times each nucleotide occurs in the string

package main
import "fmt"
func main() {

  var dna string
  countA := 0
  countC := 0
  countG := 0
  countT := 0
  flag := 0

  fmt.Println("Enter a nucleotide string : ")
  fmt.Scanln(&dna)

  for i := 0; i < len(dna); i ++ {
    if (dna[i] == 'A') {
      countA ++
    } else if (dna[i] == 'C') {
      countC ++
      } else if (dna[i] == 'G') {
        countG ++
      } else if (dna[i] == 'T') {
        countT ++
       } else {
         flag ++
       }
  }

  if flag == 0 {
  fmt.Println("A : ", countA)
  fmt.Println("C : ", countC)
  fmt.Println("G : ", countG)
  fmt.Println("T : ", countT)
} else {
  fmt.Println("Not a DNA")
}
}
