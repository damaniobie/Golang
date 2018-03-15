package main

import "fmt"

func main() {
   /* grades :=[...]int{96,54,34}
    g2 := &grades
    g2[2] = 554
    fmt.Printf("grades: %v\n",grades)
    fmt.Printf("grades2: %v",g2)*/
    ma := map[string]int{
        "cale":0,
        "cale2":300,
        "cale3":600,
    }
    a := make([]int,0,100)
    fmt.Println("len:",len(a))
    fmt.Println("cap:",cap(a))
    a = append(a,400)
    fmt.Println("len:",len(a))
    fmt.Println("cap:",cap(a))
    a = append(a,[]int{379,67,349}...) //how to add a slice to a slice
    fmt.Println(a)
    fmt.Println("len:",len(a))
    fmt.Println("cap:",cap(a))
    
    fmt.Println(ma["cale"])
    _,ok :=ma["cale"]
    fmt.Println(ok)
}
