package main

import "fmt"
import "math/rand"
       
var size int = 20
func main() {
    x := make([]int, size)
    fill_array(x)
    print_array(x)
    find_value(x,56)
    swap_indexes(x,0,19)
    //sort_array(x)
     print_array(x)
    
}


func fill_array(x[] int){
    for i:= 0;i<len(x);i++{
        x[i] = rand.Intn(100)
    }
}

func print_array(arr[] int){
    fmt.Println("Index\tValue\n-------------")
    for i:= 0;i<len(arr);i++{
        fmt.Printf("%d\t%d\n",i,arr[i])
    }
}

func find_value (arr[] int, x int) bool{
    for i:= 0;i<len(arr);i++{
        if x == arr[i]{
            fmt.Printf("%d was found at index %d\n",x, i)
            return true
        }
    }
    fmt.Println("Not found")
    return false
}

func swap_indexes(arr[] int, x int, y int){
    fmt.Println("Swapping indexes....")
    if(x < len(arr) && x >= 0 && y < len(arr) && y >= 0){
        temp := arr[x]
        arr[x] = arr[y]
        arr[y] = temp
    }
}

func sort_array(arr[]int){
     fmt.Println("Sorting....")
    for i:=0;i<len(arr)-1;i++{
        if(arr[i] < arr[i+1]){
        swap(arr[i],arr[i+1])
        }
    }
}

func swap(x int,y int){
    temp:=x
    x = y
    y = temp
}
/*func delete_index(arr[] int, index int){ this is broken for now
    if index >= 0 && index < len(arr){
        size--
    for i:= index;i<len(arr)-1;i++{
           arr[i] = arr[i+1]
    }
    
    }
}*/







