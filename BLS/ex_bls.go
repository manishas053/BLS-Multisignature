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




