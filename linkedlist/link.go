//Program to create a linked list
package main

import "fmt"

type Node struct {
  next *Node
  value int
}

type linkedList struct {
  start *Node
}

func NewNode(value int)(*Node) {
  node := &Node {
    next : nil,
    value : value,
  }
  return node
}

func newList()(*linkedList) {
  list := &linkedList {
    start : nil,
  }
  return list
}

func (list *linkedList) insertNode(value int, start *Node) {
  new := NewNode(value)
  temp := start
  if temp == nil {
    temp = new
  }else {
    for temp != nil {
      temp = temp.next
    }
    temp = new
  }
}

func (list *linkedList) printList(start *Node) {
  temp := start
  for temp != nil {
    fmt.Printf("%d -> ", temp)
    temp = temp.next
  }
}

func main() {
  var len, num int
  start := newList()
  list := newList{}
  fmt.Println("Enter the linke4d list length : ")
  fmt.Scanln(&len)
  fmt.Println("Enter the values : ")
  for i := 0; i < len; i ++ {
    fmt.Scanln(&num)
    list.insertNode(num, start)
  }
  list.printList(start)
}
