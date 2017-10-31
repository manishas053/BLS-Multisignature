//Program to determine if a sentence is a pangram.ie, checks if the sentence has every alphabet at least once.

package main
import "fmt"
import "bufio"
import "os"
func main() {
  reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    input, _ := reader.ReadString('\n')
    fmt.Println(input)

    var occ [30]int
    count:=0
    for i:=0; i<len(input); i++ {
      if (input[i] == 'a')||(input[i]=='A') {
        if occ[0] == 0 {
          count++
          occ[0] = 1
        }
      }
      if (input[i] == 'b')||(input[i]=='B') {
        if occ[1] == 0 {
          count++
          occ[1] = 1
        }
      }
      if (input[i] == 'c')||(input[i]=='C') {
        if occ[2] == 0 {
          count++
          occ[2] = 1
        }
      }
      if (input[i] == 'd')||(input[i]=='D') {
        if occ[3] == 0 {
          count++
          occ[3] = 1
        }
      }
      if (input[i] == 'e')||(input[i]=='E') {
        if occ[4] == 0 {
          count++
          occ[4] = 1
        }
      }
      if (input[i] == 'f')||(input[i]=='F') {
        if occ[5] == 0 {
          count++
          occ[5] = 1
        }
      }
      if (input[i] == 'g')||(input[i]=='G') {
        if occ[6] == 0 {
          count++
          occ[6] = 1
        }
      }
      if (input[i] == 'h')||(input[i]=='H') {
        if occ[7] == 0 {
          count++
          occ[7] = 1
        }
      }
      if (input[i] == 'i')||(input[i]=='I') {
        if occ[8] == 0 {
          count++
          occ[8] = 1
        }
      }
      if (input[i] == 'j')||(input[i]=='J') {
        if occ[9] == 0 {
          count++
          occ[9] = 1
        }
      }
      if (input[i] == 'k')||(input[i]=='K') {
        if occ[10] == 0 {
          count++
          occ[10] = 1
        }
      }
      if (input[i] == 'l')||(input[i]=='L') {
        if occ[11] == 0 {
          count++
          occ[11] = 1
        }
      }
      if (input[i] == 'm')||(input[i]=='M') {
        if occ[12] == 0 {
          count++
          occ[12] = 1
        }
      }
      if (input[i] == 'n')||(input[i]=='N') {
        if occ[13] == 0 {
          count++
          occ[13] = 1
        }
      }
      if (input[i] == 'o')||(input[i]=='O') {
        if occ[14] == 0 {
          count++
          occ[14] = 1
        }
      }
      if (input[i] == 'p')||(input[i]=='P') {
        if occ[15] == 0 {
          count++
          occ[15] = 1
        }
      }
      if (input[i] == 'q')||(input[i]=='Q') {
        if occ[16] == 0 {
          count++
          occ[16] = 1
        }
      }
      if (input[i] == 'r')||(input[i]=='R') {
        if occ[17] == 0 {
          count++
          occ[17] = 1
        }
      }
      if (input[i] == 's')||(input[i]=='S') {
        if occ[18] == 0 {
          count++
          occ[18] = 1
        }
      }
      if (input[i] == 't')||(input[i]=='T') {
        if occ[19] == 0 {
          count++
          occ[19] = 1
        }
      }
      if (input[i] == 'u')||(input[i]=='U') {
        if occ[20] == 0 {
          count++
          occ[20] = 1
        }
      }
      if (input[i] == 'v')||(input[i]=='V') {
        if occ[21] == 0 {
          count++
          occ[21] = 1
        }
      }
      if (input[i] == 'w')||(input[i]=='W') {
        if occ[22] == 0 {
          count++
          occ[22] = 1
        }
      }
      if (input[i] == 'x')||(input[i]=='X') {
        if occ[23] == 0 {
          count++
          occ[23] = 1
        }
      }
      if (input[i] == 'y')||(input[i]=='Y') {
        if occ[24] == 0 {
          count++
          occ[24] = 1
        }
      }
      if (input[i] == 'z')||(input[i]=='Z') {
        if occ[25] == 0 {
          count++
          occ[25] = 1
        }
      }
    }
    if count==26 {
      fmt.Println("Pangram")
    } else {
      fmt.Println("Not a pangram")
    }
}
