//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
//                                                                       ;;
//  Author    :  Maneesha S                                              ;;
//  Date      :  4 / 8 / 2017                                            ;;
//  Program   :  Queue implementation using linked list in golang        ;;
//                                                                       ;;
//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;


package main

import "fmt"

//defining a queue node structure
type Node struct {
  value int
  next *Node
}

//defining a queue structure
type Queue struct {
  start *Node
}

//function to create a new node
func newNode(value int) (*Node) {
  newNode := &Node{
    value : value,
    next : nil,
  }
  return newNode
}

//function to create a new queue
func newQueue() (*Queue) {
  newQueue := &Queue{
    start : nil,
  }
  return newQueue
}

//function to insert the node into queue
func (Q *Queue) insertNode(value int) {
  head := Q.start
  if head == nil {
    Q.start = newNode(value)
  } else {
    for head.next != nil {
      head = head.next
    }
    head.next = newNode(value)
  }
}

//function to print the queue elements
func (Q *Queue) printElements() {
  head := Q.start
  for head != nil {
    fmt.Printf("%d\n", head.value)
    head = head.next
  }
}

//main function
func main() {
  var limit, num int
  Q := newQueue()
  fmt.Println("Enter the number of elements : ")
  fmt.Scanln(&limit)
  fmt.Println("Enter the elements : ")
  for i := 0; i < limit; i ++ {
    fmt.Scanln(&num)
    Q.insertNode(num)
  }
  fmt.Println("Queue Elements : ")
  Q.printElements()
}
