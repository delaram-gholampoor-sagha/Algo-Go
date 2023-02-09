package main
import (
    "fmt"
    "container/heap"
)

func MedianSlidingWindow(nums []int, k int) []float64 {
    // Will store the medians
    var medians []float64 
    // Hash-map will keep track of invalid numbers
    HashMap := make(map[int]int) 
    //max heap
    SmallList := &MaxHeap{} 
    //min heap
    LargeList := &MinHeap{} 
    // Iniatlizing SmallList
    heap.Init(SmallList)
    // Iniatlizing LargeList
    heap.Init(LargeList)
    // Index of current incoming element being processed
    i := 0 
    // Initialize the SmallList heap
    for i < k{
        heap.Push(SmallList, nums[i])
        i++
    }
    // Initialize the LargeList heap 
    for j := 0; j < k/2; j++{
        heap.Push(LargeList, SmallList.Top())
        heap.Pop(SmallList)
    }
    // Start an infinite loop
    for true{
        // Get median of current window
        if k&1 == 1{ 
            medians = append(medians,float64(SmallList.Top()))
        }else{
            temp := (float64(SmallList.Top()) + float64(LargeList.Top())) * 0.5
            medians = append(medians,float64(temp))
        }
        // Break the loop if all of the elements are processed
        if i >= len(nums){
            break; 
        }
        // Outgoing element
        outNum := nums[i - k]   
        // Incoming element
        inNum := nums[i]     
        i++
        // Balance factor
        balance := 0           
        // Number `outNum` exits window
        if outNum <= SmallList.Top(){
            balance += (-1)
        }else{
            balance += 1
        }
        /* If the outgoing element is not present in the hash-map 
        store the `outNum` in the hash-map with value 1,
        otherwise increment the count of `outNum` in the hash-map.  */
        if _, ok := HashMap[outNum]; ok {  
            HashMap[outNum]++
        }else{
            HashMap[outNum] = 1
        }
        // Number `inNum` enters window
        if !SmallList.Empty() && inNum <= SmallList.Top(){
            balance++
            heap.Push(SmallList, inNum)
        }else{
            balance--
            heap.Push(LargeList, inNum)
        }
        // Re-balance SmallList 
        if balance < 0{
            heap.Push(SmallList, LargeList.Top())
            heap.Pop(LargeList)
            balance++
        }
        // Re-balance LargeList 
        if balance > 0{
            heap.Push(LargeList, SmallList.Top())
            heap.Pop(SmallList)
            balance--
        }
        // Remove invalid numbers that should be discarded from SmallList heap tops
        num, _ := HashMap[SmallList.Top()];
        for num > 0{
            HashMap[SmallList.Top()]--
            heap.Pop(SmallList)
            num, _ = HashMap[SmallList.Top()];
        }
        // Remove invalid numbers that should be discarded from LargeList heap tops
        num, _ = HashMap[LargeList.Top()];
        for !LargeList.Empty() && num > 0{
            HashMap[LargeList.Top()]--
            heap.Pop(LargeList)
            num, _ = HashMap[LargeList.Top()];
        }
    }
    // Return medians
    return medians 
    
}

func main(){
    fmt.Println("Example - 1")
    arr := []int{1,3,-1,-3,5,3,6,7}
    k := 3
    fmt.Println("Inpur: array =", arr, ", k = ",k)
    output := MedianSlidingWindow(arr,k)
    fmt.Println("Output: Medians =",output)

    fmt.Println("Example - 2")
    arr = []int{1,2}
    k = 1
    fmt.Println("Inpur: array =", arr, ", k = ",k)
    output = MedianSlidingWindow(arr,k)
    fmt.Println("Output: Medians =",output)
}