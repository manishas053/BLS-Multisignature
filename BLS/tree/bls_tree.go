//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
//                                                 ;;
//   Author   : Maneesha S                         ;;
//   Date     : 3 / 8 / 2017                       ;;
//   Program  : Binary Tree in golang              ;;
//                                                 ;;
//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

package main

import (
  "crypto/sha256"
  "fmt"
  "github.com/Nik-U/pbc"
  )

//defining a tree node structure
type treeNode struct{
  value int
  leftchild *treeNode
  rightchild *treeNode
  parent *treeNode
}

//defining a binary tree structure
type binaryTree struct{
  root *treeNode
}

//function to create a new tree node
func newNode(parent *treeNode, value int) (*treeNode) {
  newNode := &treeNode{
    value : value,
    leftchild : nil,
    rightchild : nil,
    parent : parent,
  }
  return newNode
}

//function to create an empty tree
func newTree() (*binaryTree) {
  newTree := &binaryTree{
    root : nil,
  }
  return newTree
}

//function to check if the tree node has a left child
func (tree *binaryTree) hasLeftChild(parent *treeNode) bool {
  if parent.leftchild != nil {
    return true
    } else {
      return false
    }
}

//function to check if the node has a right child
func (tree *binaryTree) hasRightChild(parent *treeNode) bool {
  if parent.rightchild != nil {
    return true
    } else {
      return false
    }
}

//function to insert the node into tree
func (tree *binaryTree) insertNode(parent *treeNode, value int) {
  if value < parent.value {
    if tree.hasLeftChild(parent) {
      tree.insertNode(parent.leftchild, value)
    } else {
      parent.leftchild = newNode(parent, value)
    }
  } else {
    if tree.hasRightChild(parent) {
      tree.insertNode(parent.rightchild, value)
    } else {
      parent.rightchild = newNode(parent, value)
    }
  }
}

//prints the binary tree
func (tree *binaryTree) printPreOrder(node *treeNode) {
  if node == nil {
    return
  }
  fmt.Printf("%d\n", node.value)
  tree.printPreOrder(node.leftchild)
  tree.printPreOrder(node.rightchild)
}

//Signature generation
func (tree *binaryTree) signatureGeneration(parent *treeNode,sharedParams string, sharedG []byte, messageChannel chan *messageData, keyChannel1 chan []byte, keyChannel2 chan []byte, keyChannel3 chan []byte, finished chan bool)) {

  if !tree.hasLeftChild(parent) && !tree.hasRightChild(parent) {
    // Node loads the system parameters
    pairing, _ := pbc.NewPairingFromString(sharedParams)
    g := pairing.NewG2().SetBytes(sharedG)

    // Generate keypairs (x, g^x)
    privKey1 := pairing.NewZr().Rand()
    pubKey1 := pairing.NewG2().PowZn(g, privKey1

    message := "some text to sign"
    h := pairing.NewG1().SetFromStringHash(message, sha256.New())

    signature1 := pairing.NewG2().PowZn(h, privKey1)
  } else {
    signatureGeneration(parent.parent, sharedParams string, sharedG []byte, messageChannel chan *messageData, keyChannel1 chan []byte, keyChannel2 chan []byte, keyChannel3 chan []byte, finished chan bool))
  }
}

// MessageData represents a signed message sent over the network
type messageData struct {
  message   string
  signature []byte
}

//main function
func main() {
  var value, limit int
  tree := newTree()
  fmt.Println("Enter number of nodes : ")
  fmt.Scanln(&limit)
  fmt.Println("Enter values : ")
  for i := 0; i < limit; i ++ {
    fmt.Scanln(&value)
    if tree.root == nil {
      tree.root = newNode(nil, value)
    }else {
      tree.insertNode(tree.root, value)
    }
  }
  fmt.Println("The binary tree is : ")
  tree.printPreOrder(tree.root)

  // The authority generates system parameters
  params := pbc.GenerateA(160, 512)
  pairing := params.NewPairing()
  g := pairing.NewG2().Rand()

  // The authority distributes system parameters params and generator g to Alice and Bob
  sharedParams := params.String()
  sharedG := g.Bytes()

  // Channel for messages
  messageChannel := make(chan *messageData)

  // Channels for public key distribution
  keyChannel1 := make(chan []byte)
  keyChannel2 := make(chan []byte)
  keyChannel3 := make(chan []byte)

  // Channel to wait until both simulations are done
  finished := make(chan bool)

  // Simulate the conversation participants
  go alice(sharedParams, sharedG, messageChannel, keyChannel1, keyChannel2, keyChannel3, finished)
  go bob(sharedParams, sharedG, messageChannel, keyChannel1,keyChannel2, keyChannel3, finished)

  // Wait for the communication to finish
  <-finished
  <-finished

}


// Alice generates a keypair and signs a message
func alice(sharedParams string, sharedG []byte, messageChannel chan *messageData, keyChannel1 chan []byte, keyChannel2 chan []byte, keyChannel3 chan []byte, finished chan bool) {
  // Alice loads the system parameters
  pairing, _ := pbc.NewPairingFromString(sharedParams)
  g := pairing.NewG2().SetBytes(sharedG)

  // Generate keypairs (x, g^x)
  privKey1 := pairing.NewZr().Rand()
  pubKey1 := pairing.NewG2().PowZn(g, privKey1)
  private2 := pairing.NewZr().Rand()
  public2 := pairing.NewG2().PowZn(g, private2)
  private3 := pairing.NewZr().Rand()
  public3 := pairing.NewG2().PowZn(g, private3)

  // Send public keys to Bob
  keyChannel1 <- pubKey1.Bytes()
  keyChannel2 <- public2.Bytes()
  keyChannel3 <- public3.Bytes()

  // Some time later, sign a message which is hashed to h, as h^x
  message := "some text to sign"
  h := pairing.NewG1().SetFromStringHash(message, sha256.New())

  signature1 := pairing.NewG2().PowZn(h, privKey1)
  signature2 := pairing.NewG2().PowZn(h, private2)
  signature3 := pairing.NewG2().PowZn(h, private3)

  // Aggregate the signatures
  aggregate1 := pairing.NewG2().Add(signature1, signature2)
  aggregate_signature := pairing.NewG2().Add(aggregate1, signature3)

  // Send the message and aggregate signature to Bob
  messageChannel <- &messageData{message: message, signature: aggregate_signature.Bytes()}

  finished <- true
}
