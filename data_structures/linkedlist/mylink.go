package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Show() {
	list := l.head
	for list != nil {
		fmt.Printf("%d -> ", list.key)
		list = list.next
	}
	fmt.Println()
}

func main() {
	l := List{}
	var num int
	var len int
	fmt.Println("Enter the length of linkedlist: ")
	fmt.Scanln(&len)
	fmt.Println("Enter values : ")
	for i := 0; i < len; i ++ {
		fmt.Scanln(&num)
		l.Insert(num)
	}
	fmt.Printf("head: %d\n", l.head.key)
	fmt.Printf("tail: %d\n", l.tail.key)
	l.Show()
}
