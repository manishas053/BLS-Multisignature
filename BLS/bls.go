// This example computes and verifies a Boneh-Lynn-Shacham signature in a simulated
// conversation between Alice and Bob.



package main

import (
  "crypto/sha256"
  "fmt"
  "github.com/Nik-U/pbc"

)

// MessageData represents a signed message sent over the network
type messageData struct {
  message   string
  signature []byte
}


func main() {
  // The authority generates system parameters
  params := pbc.GenerateA(160, 512)
  pairing := params.NewPairing()
  g := pairing.NewG2().Rand()

  // The authority distributes params and g to Alice and Bob
  sharedParams := params.String()
  sharedG := g.Bytes()

  // Channel for messages
  messageChannel := make(chan *messageData)

  // Channel for public key distribution
  keyChannel := make(chan []byte)
  keyChannel2 := make(chan []byte)
  keyChannel3 := make(chan []byte)

  // Channel to wait until both simulations are done
  finished := make(chan bool)

  // Simulate the conversation participants
  go alice(sharedParams, sharedG, messageChannel, keyChannel, keyChannel2, keyChannel3, finished)
  go bob(sharedParams, sharedG, messageChannel, keyChannel,keyChannel2, keyChannel3, finished)

  // Wait for the communication to finish
  <-finished
  <-finished

}


// Alice generates a keypair and signs a message
func alice(sharedParams string, sharedG []byte, messageChannel chan *messageData, keyChannel chan []byte, keyChannel2 chan []byte, keyChannel3 chan []byte, finished chan bool) {
  // Alice loads the system parameters
  pairing, _ := pbc.NewPairingFromString(sharedParams)
  g := pairing.NewG2().SetBytes(sharedG)

  // Generate keypairs (x, g^x)
  privKey := pairing.NewZr().Rand()
  pubKey := pairing.NewG2().PowZn(g, privKey)
  private2 := pairing.NewZr().Rand()
  public2 := pairing.NewG2().PowZn(g, private2)
  private3 := pairing.NewZr().Rand()
  public3 := pairing.NewG2().PowZn(g, private3)

  // Send public keys to Bob
  keyChannel <- pubKey.Bytes()
  keyChannel2 <- public2.Bytes()
  keyChannel3 <- public3.Bytes()

  // Some time later, sign a message, hashed to h, as h^x
  message := "some text to sign"
  h := pairing.NewG1().SetFromStringHash(message, sha256.New())

  signature1 := pairing.NewG2().PowZn(h, privKey)
  signature2 := pairing.NewG2().PowZn(h, private2)
  signature3 := pairing.NewG2().PowZn(h, private3)

  // Aggregate the signatures
  aggregate1 := pairing.NewG2().Add(signature1, signature2)
  aggregate_signature := pairing.NewG2().Add(aggregate1, signature3)

  // Send the message and signature to Bob
  messageChannel <- &messageData{message: message, signature: aggregate_signature.Bytes()}

  finished <- true
}


// Bob verifies a message received from Alice
func bob(sharedParams string, sharedG []byte, messageChannel chan *messageData, keyChannel chan []byte, keyChannel2 chan []byte, keyChannel3 chan []byte, finished chan bool) {
  // Bob loads the system parameters
  pairing, _ := pbc.NewPairingFromString(sharedParams)
  g := pairing.NewG2().SetBytes(sharedG)

  // Bob receives Alice's public keys
  pubKey := pairing.NewG2().SetBytes(<-keyChannel)
  public2 := pairing.NewG2().SetBytes(<-keyChannel2)
  public3 := pairing.NewG2().SetBytes(<-keyChannel3)

  // Some time later, Bob receives a message to verify
  data := <-messageChannel
  signature := pairing.NewG1().SetBytes(data.signature)

  // To verify, Bob checks that e(h,g^x)=e(sig,g)
  h := pairing.NewG1().SetFromStringHash(data.message, sha256.New())
  temp1 := pairing.NewGT().Pair(h, pubKey)
  temp3 := pairing.NewGT().Pair(h, public2)
  temp4 := pairing.NewGT().Pair(h, public3)

  temp5 := pairing.NewGT().Add(temp1, temp3)
  temp := pairing.NewGT().Add(temp4, temp5)

  temp2 := pairing.NewGT().Pair(signature, g)
  if !temp.Equals(temp2) {
      fmt.Println("*BUG* Signature check failed *BUG*")
  } else {
      fmt.Println("Signature verified correctly")
  }

  finished <- true
}
