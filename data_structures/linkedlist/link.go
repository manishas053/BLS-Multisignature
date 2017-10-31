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

func (newlist *linkedList) insertNode(value int, start *Node) {
  //fmt.Println("hi")
  new := NewNode(value)
  temp := start
  fmt.Println(temp)
  if temp == nil {
    start = new
    fmt.Println(start)
  }else {
    for temp != nil {
      //fmt.Println("hi")
      temp = temp.next
    }
    temp = new
  }
}

func (newlist *linkedList) printList(start *Node) {
  temp := start
  for temp != nil {
    fmt.Printf("%d -> ", temp)
    temp = temp.next
  }
}

func main() {
  var len, num int
  newlist := newList()
  //list := linkedList{}
  fmt.Println("Enter the linked list length : ")
  fmt.Scanln(&len)
  fmt.Println("Enter the values : ")
  for i := 0; i < len; i ++ {
    fmt.Scanln(&num)
    newlist.insertNode(num, newlist.start)
  }
  newlist.printList(newlist.start)
}
