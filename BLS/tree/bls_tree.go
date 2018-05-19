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
  "sync"
  "time"
  "github.com/Nik-U/pbc"
  )

//defining a tree node structure
type treeNode struct{
  value int
  leftchild *treeNode
  rightchild *treeNode
  parent *treeNode
  aggregate []byte
}

//defining a binary tree structure
type binaryTree struct{
  root *treeNode
}

// MessageData represents a signed message sent over the network
type messageData struct {
  message   string
  signature []byte
}

//function to create a new tree node
func newNode(parent *treeNode, value int) (*treeNode) {
  newNode := &treeNode{
    value : value,
    leftchild : nil,
    rightchild : nil,
    parent : parent,
    aggregate : nil,
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

//Signature aggregation
func (tree *binaryTree) aggregation(parent *treeNode, signature1 []byte, sharedParams string, sharedG []byte ) {
  pairing, _ := pbc.NewPairingFromString(sharedParams)
  // Generate keypairs (x, g^x)
  privKey1 := pairing.NewZr().Rand()
  message := "some text to sign"
  h := pairing.NewG1().SetFromStringHash(message, sha256.New())

  signature := pairing.NewG2().PowZn(h, privKey1)

  //parent.aggregate = parent.aggregate + number
  if (tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    sign1 := pairing.NewG2().SetBytes(parent.leftchild.aggregate)
    sign2 := pairing.NewG2().SetBytes(parent.rightchild.aggregate)
    aggregate1 := pairing.NewG2().Add(sign1, sign2)
    aggregate2 := pairing.NewG2().Add(aggregate1, signature)
    parent.aggregate = aggregate2.Bytes()
    fmt.Printf("%d : %d\n\n", parent.value, parent.aggregate)

  } else if (tree.hasLeftChild(parent) && !tree.hasRightChild(parent)) {
    sign1 := pairing.NewG2().SetBytes(parent.leftchild.aggregate)
    aggregate1 := pairing.NewG2().Add(sign1, signature)
    parent.aggregate = aggregate1.Bytes()
    fmt.Printf("%d : %d\n\n", parent.value, parent.aggregate)
  } else if (!tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    sign1 := pairing.NewG2().SetBytes(parent.rightchild.aggregate)
    aggregate1 := pairing.NewG2().Add(sign1, signature)
    parent.aggregate = aggregate1.Bytes()
    fmt.Printf("%d : %d\n\n", parent.value, parent.aggregate)
  }
  tree.aggregation(parent.parent, parent.aggregate, sharedParams, sharedG)
}

//Signature generation
func (tree *binaryTree) signatureGeneration(parent *treeNode,sharedParams string, sharedG []byte ) {
  var wg sync.WaitGroup
  wg.Add(2)
  if (tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    go tree.signatureGeneration(parent.leftchild, sharedParams, sharedG)
    go tree.signatureGeneration(parent.rightchild, sharedParams, sharedG)
    wg.Wait()
    time.Sleep(1000 * time.Millisecond)
  } else if (tree.hasLeftChild(parent) && !tree.hasRightChild(parent)) {
    tree.signatureGeneration(parent.leftchild, sharedParams, sharedG)
  } else if (!tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    tree.signatureGeneration(parent.rightchild, sharedParams, sharedG)
  } else {
    // Node loads the system parameters
    pairing, _ := pbc.NewPairingFromString(sharedParams)
    //g := pairing.NewG2().SetBytes(sharedG)

    // Generate keypairs (x, g^x)
    privKey1 := pairing.NewZr().Rand()
  //  pubKey1 := pairing.NewG2().PowZn(g, privKey1)

    message := "some text to sign"
    h := pairing.NewG1().SetFromStringHash(message, sha256.New())

    signature1 := pairing.NewG2().PowZn(h, privKey1)
    parent.aggregate = signature1.Bytes()
    tree.aggregation(parent.parent, parent.aggregate, sharedParams, sharedG)
    time.Sleep(500 * time.Millisecond)
  }
  defer wg.Done()
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

  fmt.Println("\nAggregate Signatures :")
  tree.signatureGeneration(tree.root, sharedParams, sharedG)
}
