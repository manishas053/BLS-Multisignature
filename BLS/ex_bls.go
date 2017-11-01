package main

import (
    "crypto/sha256"
    "fmt"
    
    "pbc"
)

// messageData represents a signed message sent over the network
type messageData struct {
    message   string
    signature []byte
}


func main() {
    // The authority generates system parameters
    
    pairing := params.NewPairing()
    g := pairing.NewG2().Rand()

    // The authority distributes params and g to signer and verifier
    sharedParams := params.String()
    sharedG := g.Bytes()

    // Channel for messages. Normally this would be a network connection.
    messageChannel := make(chan *messageData)

    // Channel for public key distribution. 
    keyChannel := make(chan []byte)

    // Channel to wait until both simulations are done
    finished := make(chan bool)

    // Simulate the conversation participants
    go signer(sharedParams, sharedG, messageChannel, keyChannel, finished)
    go verifier(sharedParams, sharedG, messageChannel, keyChannel, finished)

    // Wait for the communication to finish
    <-finished
    <-finished

}

// Signer generates a keypair and signs a message
func signer(sharedParams string, sharedG []byte, messageChannel chan *messageData, keyChannel chan []byte, finished chan bool) {
    // Alice loads the system parameters
    pairing, _ := pbc.NewPairingFromString(sharedParams)
    g := pairing.NewG2().SetBytes(sharedG)

    // Generate keypair (x, g^x)
    privKey := pairing.NewZr().Rand()
    pubKey := pairing.NewG2().PowZn(g, privKey)

    // Send public key to verifier
    keyChannel <- pubKey.Bytes()

    // Some time later, sign a message, hashed to h, as h^x
    message := "hello world"
    h := pairing.NewG1().SetFromStringHash(message, sha256.New())
    signature := pairing.NewG2().PowZn(h, privKey)

    // Send the message and signature to verifier
    messageChannel <- &messageData{message: message, signature: signature.Bytes()}

    finished <- true
}


