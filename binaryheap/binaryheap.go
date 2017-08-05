package main

import "fmt"


func heapify(i int, heap []int) []int {
  if i/2 < 0 {
    return heap
  }
  if heap[i] < heap[i/2] {
    heap[i], heap[i/2] = swap(heap[i], heap[i/2])
    heap = heapify(i/2, heap)
  }
  return heap
}

func swap(num1, num2 int) (int, int) {
  temp := num1
  num1 = num2
  num2 = temp
  return num1, num2

}

func main() {
  var len, num, option int
  fmt.Printf("Choose any option\n1. Create a heap\n2. Insert a node\n3. Delete a node\n")
  fmt.Scanln(&option)
  switch option {
  case 1:
    fmt.Println("Enter the length : ")
    fmt.Scanln(&len)
    heap := make([]int, len+1)
    fmt.Println("Enter elements : ")
    for i := 1; i <= len; i ++ {
    fmt.Scanln(&num)
    heap[i] = num
    heap = heapify(i, heap)
}
fmt.Println(heap)
  }
}
