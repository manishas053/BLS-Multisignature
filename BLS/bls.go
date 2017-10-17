package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
  "bufio"
  "os"
  "crypto/sha256"
)

//function to generate a random number between a given range
func random_gen(min,max int) int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(max - min) + min
}

//main function
func main() {
  var r int
  fmt.Println("Large prime r : ")
  fmt.Scanln(&r)

  //generator of the group G
  generator := random_gen(1, r - 1)
  fmt.Println("generator of the group  :", generator)

  //private key
  private := random_gen(1, generator)
  fmt.Println("private key \t\t:", private)

  //public key
  public := math.Pow(float64(generator), float64(private))
  fmt.Println("public key \t\t:", public)

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter message   \t: ")
  message, _ := reader.ReadString('\n')

  //sha256 hash of the message
  hash := sha256.Sum256([]byte(message))
  fmt.Printf("hash of the message\t: %X\n", hash)

  //signature on the hash
  signature := math.Pow(float64(hash), float64(private))
  fmt.Println(signature)
}
