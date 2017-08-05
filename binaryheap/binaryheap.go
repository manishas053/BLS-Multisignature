//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
//                                                                          ;;
//  Author  :  Maneesha S                                                   ;;
//  Date    :  5 / 8 / 2017                                                 ;;
//  Program : Binary heap implementation in golang                          ;;
//                                                                          ;;
//;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;


package main

import (
  "fmt"
  "os"
)

//function to check the heap order property while deleting a node
func percDown(pos int, heap []int) []int {
  if pos < 0 {
    return heap
  }
  if ((2 * pos) < len(heap)) || (((2 * pos) + 1) < len(heap)) {
    if ((heap[2 * pos] == 0) && (heap[(2 * pos) + 1] == 0)) {
      return heap
      } else {
        if ((heap[2 * pos] < heap[(2 * pos) + 1] )) {
          if heap[pos] > heap[2 * pos] {
            heap[pos], heap[2 * pos] = swap(heap[pos], heap[2 * pos])
          }
        } else {
          if ((heap[((2 * pos) + 1)] != 0) && ((2 * pos) + 1) < len(heap)) {
            if heap[pos] > heap[(2 * pos) + 1] {
              heap[pos], heap[(2 * pos) + 1] = swap(heap[pos], heap[(2 * pos) + 1])
            }
          }
        }
      }
  }
return heap
}

//function to delete a node from the heap
func delete(num int, heap []int) []int {
  var i, pos int
  for i := 1; i < len(heap); i ++ {
    if heap[i] == num {
      pos = i
    }
  }
  i = len(heap)
  for heap[i - 1] == 0 {
    i --
  }
  heap[pos] = heap[i - 1]
  heap[i - 1] = 0
  heap = percDown(pos, heap)
  return heap
}

//function to check the heap order property while inserting a node
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

//function to swap two numbers
func swap(num1, num2 int) (int, int) {
  temp := num1
  num1 = num2
  num2 = temp
  return num1, num2

}

//function to choose the operation
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

  case 3:
    fmt.Println("Enter the element to be deleted : ")
    fmt.Scanln(&num)
    heap = delete(num, heap)
    fmt.Println(heap)
  }
}

//if user want to continue then calls another function called sample else the program exits
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

//main function
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
