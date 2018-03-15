package main

import (
    "fmt"
    "sort"
      )
  
/* Merge two arrays, handles sorted and unsorted, as well as arrays of different sizes*/

func main() {
    x := []int{0,2,5,8,15,19,28,47,59,88,99}
    y := []int{1,4,6,16,26,33,47,68,103,115}
    unsorted_arr1 := []int{45,87,23,10,778,2,98,168,453,0,345,3}
    unsorted_arr2 := []int{667,34,65,89,997,554,34,89,10,167,44,67,2,8,1000,5}
    
    fmt.Println(mergeArr(unsorted_arr1,unsorted_arr2))
    fmt.Println(mergeArr(x,y))
}

func mergeArr (arr1, arr2 []int)[]int{
    merge := make([]int,len(arr1)+len(arr2))
    sort.Ints(arr1)
    sort.Ints(arr2)
    index1,index2,merge_index := 0,0,0
    
    for merge_index < len(merge){
        exhausted1 := index1 >= len(arr1)
        exhausted2 := index2 >= len(arr2)
        
        if !exhausted1 && (exhausted2 || (arr1[index1] < arr2[index2])){
            merge[merge_index] = arr1[index1]
            index1++
        }else{
             merge[merge_index] = arr2[index2]
            index2++
        }
        merge_index++
    }
    
    return merge
}
