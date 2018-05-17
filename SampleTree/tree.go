//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
//                                                 ;;
//   Author   : Maneesha S                         ;;
//   Date     : 16 / 5 / 2018                       ;;
//   Program  : Binary Tree in golang              ;;
//                                                 ;;
//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

package main

import (
  "fmt"
  "sync"
)

//defining a tree node structure
type treeNode struct{
  value int
  leftchild *treeNode
  rightchild *treeNode
  parent *treeNode
  aggregate int
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
    aggregate : 0,
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

//aggregation
func (tree *binaryTree) aggregation(parent *treeNode, number int) {
  //defer wg.Done()
  //parent.aggregate = parent.aggregate + number
  if (tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    parent.aggregate = parent.leftchild.aggregate + parent.rightchild.aggregate
    fmt.Printf("%d : %d\n", parent.value, parent.aggregate)
  } else if (tree.hasLeftChild(parent) && !tree.hasRightChild(parent)) {
    parent.aggregate = parent.leftchild.aggregate
    fmt.Printf("%d : %d\n", parent.value, parent.aggregate)
  } else if (!tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    parent.aggregate = parent.rightchild.aggregate
    fmt.Printf("%d : %d\n", parent.value, parent.aggregate)
  }
  tree.aggregation(parent.parent, parent.aggregate)
}

//leaf child
func (tree *binaryTree) passValue(parent *treeNode, number int) {
  var wg sync.WaitGroup
  wg.Add(2) 
  if (tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    go tree.passValue(parent.leftchild, number)
    go tree.passValue(parent.rightchild, number)
    wg.Wait()
    //time.Sleep(1000 * time.Millisecond)
  } else if (tree.hasLeftChild(parent) && !tree.hasRightChild(parent)) {
    tree.passValue(parent.leftchild, number)
  } else if (!tree.hasLeftChild(parent) && tree.hasRightChild(parent)) {
    tree.passValue(parent.rightchild, number)
  } else {
    parent.aggregate = number
    fmt.Printf("%d : %d\n", parent.value, parent.aggregate)
    //time.Sleep(500 * time.Millisecond)
    tree.aggregation(parent.parent, parent.aggregate)
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

  number := 2
  fmt.Println("PASSVALUE FUNC")
  tree.passValue(tree.root, number)

  fmt.Printf("Root %d : $d\n", tree.root, tree.root.aggregate)

}
