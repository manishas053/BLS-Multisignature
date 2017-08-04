package main

import "fmt"

type Node struct {
  value int
  next *Node
  prev *Node
}

type Stack struct {
  start *Node
}

func newNode(node *Node, value int) (*Node) {
  newNode := &Node{
    value : value,
    next : nil,
    prev : node,
  }
  return newNode
}

func newStack() (*Stack) {
  newStack := &Stack{
    start : nil,
  }
  return newStack
}

func (stack *Stack) push(value int) {
  head := stack.start
  if head == nil {
    stack.start = newNode(nil, value)
  } else {
    for head.next != nil {
      head = head.next
    }
    head.next = newNode(head, value)
  }
}

func (stack *Stack) pop() {
  head := stack.start
  if head == nil {
    return
  } else {
    for head.next != nil {
      head = head.next
    }
    temp := head
    for temp != nil {
      fmt.Printf("%d\n", temp.value)
      temp = temp.prev
    }
  }
}

func main() {
  var limit, num int
  stack := newStack()
  fmt.Println("Enter the number of elements : ")
  fmt.Scanln(&limit)
  fmt.Println("Enter the elements : ")
  for i := 0; i < limit; i ++ {
    fmt.Scanln(&num)
    stack.push(num)
  }
  fmt.Println("Elements in the stack are : ")
  stack.pop()
}
