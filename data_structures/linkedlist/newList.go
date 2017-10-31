///////////////////////////////////////////////////////////////////////
//                                                                  //
//   Author   : Maneesha S                                         //
//   Date     : 3 / 8 / 2017                                      //
//   Program  : Linked List in golang                            //
//                                                              //
/////////////////////////////////////////////////////////////////

package main

import "fmt"

type Node struct {
  next *Node
  value int
}

type linkedList struct {
  start *Node
}

//function to create an object for a single node
func newNode(value int) (*Node) {
  newNode := &Node{
    next : nil,
    value : value,
  }
  return newNode
}

//function to create an empty linked list
func newList() (*linkedList) {
  newList := &linkedList{
    start : nil,
    }
  return newList
}

//function to insert a new node to the linked list
func (list *linkedList) insertNode(value int) {
  node := newNode(value)
  head := list.start
  if head == nil {
    list.start = node
  }else {
    for head.next != nil {
      head = head.next
    }
    head.next = node
  }
}

//prints the linked list
func (list *linkedList) printNode(start *Node) {
  head := list.start
  for head != nil {
    fmt.Printf("%d -> ", head.value)
    head = head.next
    }
  fmt.Println()
}

func main() {
  var length, num int
  list := newList()
  fmt.Println("Enter linked list length : ")
  fmt.Scanln(&length)
  fmt.Println("Enter the values : ")
  for i := 0; i < length; i ++ {
    fmt.Scanln(&num)
    list.insertNode(num)
  }
  list.printNode(list.start)
}
