package main
import "fmt"

type Node struct {
  value int
  leftchild *Node
  rightchild *Node
  parent *Node
  leftheight int
  rightheight int
}

type Tree struct {
  root *Node
}

func newNode(parent *Node, value int) *Node {
  newNode := &Node {
    value : value,
    leftchild : nil,
    rightchild : nil,
    parent : parent,
    leftheight : 0,
    rightheight : 0,
  }
  return newNode
}

func newTree() *Tree {
  newTree := &Tree{
    root : nil,
  }
  return newTree
}

func (tree *Tree) leftHeight(parent *Node) int {
  parent.leftheight ++
  if parent.leftheight - parent.rightheight == 2 {
    if parent.leftchild.leftheight - parent.leftchild.rightheight == 1 {
      parent = tree.leftLeftRotate(parent)
    } else {
      parent = tree.leftRightRotate(parent)
    }
  }
  if parent != tree.root {
    if parent == parent.parent.leftchild {
      tree.leftHeight(parent.parent)
    } else {
      tree.rightHeight(parent.parent)
    }
  }
  return parent.leftheight
}

func (tree *Tree) rightHeight(parent *Node) int {
  parent.rightheight ++
  if parent.rightheight - parent.leftheight == 2 {
    if parent.rightchild.rightheight - parent.rightchild.leftheight == 1 {
      parent = tree.rightRightRotate(parent)
    } else {
      parent = tree.rightLeftRotate(parent)
    }
  }

  if parent != tree.root {
    if parent == parent.parent.leftchild {
      tree.leftHeight(parent.parent)
    } else {
      tree.rightHeight(parent.parent)
    }
  }
  return parent.rightheight
}

func (tree *Tree) leftLeftRotate(parent *Node) (*Node) {
  temp := parent
  parent = parent.leftchild
  parent.rightchild = temp
  return parent
}

func (tree *Tree) leftRightRotate(parent *Node) (*Node) {
  temp := parent.leftchild
  parent.leftchild = parent.leftchild.rightchild
  parent.leftchild.leftchild = temp
  tree.leftLeftRotate(parent)
  return parent
}

func (tree *Tree) rightRightRotate(parent *Node) (*Node) {
  temp := parent
  parent = parent.rightchild
  parent.leftchild = temp
  return parent
}

func (tree *Tree) rightLeftRotate(parent *Node) (*Node) {
  temp := parent.rightchild
  parent.rightchild = parent.rightchild.leftchild
  parent.rightchild.rightchild = temp
  tree.rightRightRotate(parent)
  return parent
}

func (tree *Tree) hasLC(parent *Node) bool {
  if parent.leftchild != nil {
    return true
    } else {
      return false
    }
}

//function to check if the node has a right child
func (tree *Tree) hasRC(parent *Node) bool {
  if parent.rightchild != nil {
    return true
    } else {
      return false
    }
}

func (tree *Tree) printInOrder(node *Node) {
  if node == nil {
    return
  }
  tree.printInOrder(node.leftchild)
  fmt.Printf("%d\n", node.value)
  tree.printInOrder(node.rightchild)
}

func (tree *Tree) insertNode(parent *Node, value int) {
  if value < parent.value {
    if tree.hasLC(parent) {
      tree.insertNode(parent.leftchild, value)
    } else {
      parent.leftchild = newNode(parent, value)
      tree.leftHeight(parent)
    }

  } else {
    if tree.hasRC(parent) {
      tree.insertNode(parent.rightchild, value)
    } else {
      parent.rightchild = newNode(parent, value)
      tree.rightHeight(parent)
    }
  }
}

func main() {
  var limit, num int
  tree := newTree()
  fmt.Println("Enter the limit :")
  fmt.Scanln(&limit)
  fmt.Println("Enter the numbers :")
  for i := 0; i < limit; i ++ {
    fmt.Scanln(&num)
    if tree.root == nil {
      tree.root = newNode(nil, num)
    } else {
      tree.insertNode(tree.root, num)
    }
  }
  fmt.Println("AVL Tree is : ")
  tree.printInOrder(tree.root)
}
