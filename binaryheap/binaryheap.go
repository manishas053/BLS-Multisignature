package main

import (
  "fmt"
  "os"
)


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
  var len int
  fmt.Println("Enter the maximum length : ")
  fmt.Scanln(&len)
  heap := make([]int, len+1)
  for {
    fmt.Println("Want to continue (y or n) ?")
    option(heap)
  }
}

func option(heap []int) {
  var char string
  fmt.Scanln(&char)
  switch char {
  case "y":
    sample(heap)
    return
  case "n":
    os.Exit(2)
  default:
    fmt.Println("Invalid input : Press y to continue or n to exit ")
  }
}

func sample(heap []int) {
  var limit, num, option int
  fmt.Printf("Choose any option\n1. Create a heap\n2. Insert a node\n3. Delete a node\n")
  fmt.Scanln(&option)
  switch option {
  case 1:
    fmt.Println("Enter the number of elements : ")
    fmt.Scanln(&limit)
    fmt.Println("Enter elements : ")
    for i := 1; i <= limit; i ++ {
      fmt.Scanln(&num)
      heap[i] = num
      heap = heapify(i, heap)
    }
    fmt.Println(heap)
    return

  case 2:
    fmt.Println("Enter the element to be inserted : ")
    fmt.Scanln(&num)
    i := 1
    for heap[i] != 0 {
      i ++
    }
    heap[i] = num
    heap = heapify(i, heap)
    fmt.Println(heap)
  }
}
