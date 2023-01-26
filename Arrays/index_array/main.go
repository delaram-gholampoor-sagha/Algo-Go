package main

import("fmt")

func indexArray(arr []int, size int) { 
  for i := 0; i < size; i++ {
    curr := i
    value := -1
/* swaps to move elements in the proper position. */ 
    for arr[curr] != -1 && arr[curr] != curr {
      
        temp := arr[curr]
        arr[curr] = value
        value = temp
        curr = temp
    }
/* check if some swaps happened. */ 
        if value != -1 {
            arr[curr] = value
        } 
  } 
}
/* Testing code */
func main() {
    arr := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
    size := len(arr)
    indexArray(arr, size)
    fmt.Println(arr)
}