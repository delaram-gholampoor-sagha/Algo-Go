package main
import "fmt"

func SmallestPositiveMissingNumber2(arr []int, size int) int { 
  hs := make(map[int]int)
  for i := 0; i < size; i++ {
        hs[arr[i]] = 1
  }
    for i := 1; i < size+1; i++ {
        _, ok := hs[i]
        if ok == false {
            return i
        } 
    }

return -1 
}
//Testing code
func main() {
    arr := []int{8, 5, 6, 1, 9, 11, 2, 7, 4, 10}
    size := len(arr)
    fmt.Println("Missing Number :", SmallestPositiveMissingNumber2(arr,size))

}