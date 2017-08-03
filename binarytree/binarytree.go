///////////////////////////////////////////////////////
//                                                  //
//   Author   : Maneesha S                         //
//   Date     : 3 / 8 / 2017                      //
//   Program  : Binary Tree in golang            //
//                                              //
/////////////////////////////////////////////////

package main

import "fmt"

type treeNode struct{
  value int
  leftchild *treeNode
  rightchild *treeNode
  parent *treeNode
}

type binaryTree struct{
  root *treeNode
}


//function to create an object for a node
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


//function to check if a node has leftchild
func (tree *binaryTree) hasLeftChild(parent *treeNode) bool {
  if parent.leftchild != nil {
    return true
    }else {
      return false
      }
}


//function to check if a node has right child
func (tree *binaryTree) hasRightChild(parent *treeNode) bool {
  if parent.rightchild != nil {
    return true
    }else {
      return false
      }
}


//function to insert node into tree
func (tree *binaryTree) insertNode(parent *treeNode, value int) {
  if value < parent.value {
    if tree.hasLeftChild(parent) {
      tree.insertNode(parent.leftchild, value)
    }else {
      parent.leftchild = newNode(parent, value)
    }
  }else {
    if tree.hasRightChild(parent) {
      tree.insertNode(parent.rightchild, value)
    }else {
      parent.rightchild = newNode(parent, value)
    }
  }
}


//prints the binary tree
func (tree *binaryTree) printInorder(node *treeNode) {
  if node == nil {
    return
  }
  fmt.Printf("%d\n", node.value)
  tree.printInorder(node.leftchild)
  tree.printInorder(node.rightchild)
}

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
  tree.printInorder(tree.root)
}
