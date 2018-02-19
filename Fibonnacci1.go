package main

import "fmt"

/* Fibonacci program using recursion. The max index for this is the 43rd Fibonnacci number*/
func fib(x int ) int {
    if !(x<1){
    if x == 1{
        return 0
    }else if x == 2 || x == 3{
        return 1
    }else{
        return fib(x-1) + fib(x - 2)
    }
    }else{
        fmt.Println("ERROR - MUST BE A POSITIVE NUMBER - RETURNING ERROR CODE '-1'")
        return -1 //error
    }
}
func main() {
    fib_variable := 43 //the highest fib index it'll take.. working on this
    x := fib(fib_variable)
    fmt.Printf("The %dth fibonacci number is %d",fib_variable, x)
    

}
