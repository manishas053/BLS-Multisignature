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

func newNode(parent, value int) (*treeNode) {
  newNode := &treeNode{
    value : value,
    leftchild : nil,
    rightchild : nil,
    parent : parent,
  }
  return newNode
}

func newTree() (*binaryTree) {
  newTree := &binaryTree{
    root : nil,
  }
  return newTree
}

func (tree *binaryTree) hasLeftChild(parent *treeNode) bool {
  if parent.leftchild != nil {
    return true
    }else {
      return false
      }
}

func (tree *binaryTree) hasRightChild(parent *treeNode) bool {
  if parent.rightchild != nil {
    return true
    }else {
      return false
      }
}

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

func printInorder(node *treeNode) {
  if node == nil {
    return
  }
  fmt.Println("%d\n", node.value)
  printInorder(node.leftchild)
  printInorder(node.rightchild)
}

func main() {
  var value int
  tree := newTree()
  fmt.Println("Enter values : ")
  fmt.Scanln(&value)
  if tree.root == nil {
    tree.root = value
  }else {
    tree.insertNode(tree.root, value)
  }
  tree.printInorder(tree.root)
}
