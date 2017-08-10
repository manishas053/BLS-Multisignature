//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
//                                                                        ;;
//    Author     :  Maneesha S                                            ;;
//    Date       :  8 / 8 / 2017                                          ;;
//    Program    :  Implemention of AVL tree in golang                    ;;
//                                                                        ;;
//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;


package main
import "fmt"

//Defining an avl tree node structure
type Node struct {
  value int
  leftchild *Node
  rightchild *Node
  parent *Node
  leftheight int
  rightheight int
}

//Defining an avl tree structure
type Tree struct {
  root *Node
}

//Function to create an avl node
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

//Function to create a new avl tree
func newTree() *Tree {
  newTree := &Tree{
    root : nil,
  }
  return newTree
}

//Function to increment leftheight of the parent and do left left or left right rotation if required
func (tree *Tree) leftHeight(node *Node) (*Node) {
  node.leftheight ++
  if node != tree.root {
    temp := node
    node := node.parent
    if node.leftchild == temp {
      node.leftheight ++
    } else {
      node.rightheight ++
    }
    if node.leftheight - node.rightheight == 2 {
      if temp.leftheight - temp.rightheight == 1 {
        parent := tree.leftLeftRotate(node)
        if node == tree.root {
          tree.root = parent
        }
      } else {
        parent := tree.leftRightRotate(node)
        if node == tree.root {
          tree.root = parent
        }
      }
    }
  }
  return tree.root
}

//Function to increment right height of the parent and do right right or right left rotation if required
func (tree *Tree) rightHeight(node *Node) (*Node) {
  node.rightheight ++
  if node != tree.root {
    if node == node.parent.leftchild {
      node.parent.leftheight ++
    } else {
      node.parent.rightheight ++
    }
    if node.parent.rightheight - node.parent.leftheight == 2 {
      if node.rightheight - node.leftheight == 1 {
        parent := tree.rightRightRotate(node.parent)
        if node.parent == tree.root {
          tree.root = parent
        }
      } else {
        parent := tree.rightLeftRotate(node.parent)
        if node.parent == tree.root {
          tree.root = parent
        }
      }
    }
  }
  return tree.root
}

//Function for left left rotation
func (tree *Tree) leftLeftRotate(parent *Node) (*Node) {
  temp := parent
  if parent != tree.root {
    parent.parent.leftchild = parent.leftchild
  }
  temp2 := parent.rightchild
  parent = parent.leftchild
  parent.rightchild = temp
  temp.parent = parent
  parent.rightchild.leftchild = temp2
  parent.rightheight = 1
  parent.rightchild.leftheight = 0
  return parent
}

//Function for left right rotation
func (tree *Tree) leftRightRotate(parent *Node) (*Node) {
  temp := parent.leftchild
  parent.leftchild = parent.leftchild.rightchild
  parent.leftchild.leftchild = temp
  parent.leftchild.leftheight = 1
  parent.leftchild.leftchild.rightheight = 0
  tree.leftLeftRotate(parent)
  return parent
}

//Function for right right rotation
func (tree *Tree) rightRightRotate(parent *Node) (*Node) {
  temp := parent
  if parent != tree.root {
    parent.parent.rightchild = parent.rightchild
  }
  temp2 := parent.leftchild
  parent = parent.rightchild
  parent.leftchild = temp
  temp.parent = parent
  parent.leftchild.rightchild = temp2
  parent.leftheight = 1
  parent.leftchild.rightheight = 0
  return parent
}

//Function for right left rotation
func (tree *Tree) rightLeftRotate(parent *Node) (*Node) {
  temp := parent.rightchild
  parent.rightchild = parent.rightchild.leftchild
  parent.rightchild.rightheight = 1
  parent.rightchild.rightchild = temp
  parent.rightchild.rightchild.leftheight = 0
  tree.rightRightRotate(parent)
  return parent
}

//Checks if the node has a left child
func (tree *Tree) hasLC(parent *Node) bool {
  if parent.leftchild != nil {
    return true
    } else {
      return false
    }
}

//Checks if the node has a right child
func (tree *Tree) hasRC(parent *Node) bool {
  if parent.rightchild != nil {
    return true
    } else {
      return false
    }
}

//Prints tree nodes in inorder
func (tree *Tree) printInOrder(node *Node) {
  if node == nil {
    return
  }
  tree.printInOrder(node.leftchild)
  fmt.Printf("%d\n", node.value)
  tree.printInOrder(node.rightchild)
}

//Function to insert the node into avl tree
func (tree *Tree) insertNode(parent *Node, value int) (*Node) {
  if value < parent.value {
    if tree.hasLC(parent) {
      tree.insertNode(parent.leftchild, value)
    } else {
      parent.leftchild = newNode(parent, value)
      parent = tree.leftHeight(parent)
    }

  } else {
    if tree.hasRC(parent) {
      tree.insertNode(parent.rightchild, value)
    } else {
      parent.rightchild = newNode(parent, value)
      parent = tree.rightHeight(parent)
    }
  }
  return tree.root
}

//Main function
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
      tree.root = tree.insertNode(tree.root, num)
    }
  }
  fmt.Println("AVL Tree is : ")
  tree.printInOrder(tree.root)
}
