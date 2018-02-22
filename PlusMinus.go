package main
import "fmt"

/*
Given an array of integers, calculate which fraction of its elements are positive, which fraction of its elements are negative, 
and which fraction of its elements are zeroes, respectively. Print the decimal value of each fraction on a new line.

*/


func plus_minus(arr []int){
    var num_pos,num_neg,num_0 float64
    num_pos,num_neg,num_0 = 0.0,0.0,0.0
    for i := 0;i<len(arr);i++{
        if arr[i] > 0{
            num_pos++
        }else if arr[i] < 0{
            num_neg++
        }else{
            num_0++
        }
    }
    fmt.Printf("%f\n%f\n%f",(num_pos / float64(len(arr))),(num_neg / float64(len(arr))),(num_0 / float64(len(arr))))
}

func main() {
    //size := 6
    array := []int{-4,3,-9,0,4,1}
    plus_minus(array)
    
    
}
