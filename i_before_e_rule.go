package main

import "fmt"

func main() {
    //fmt.Printf("Hello, World!")
    fmt.Println(check2("dog"))
}

func check(s string) (v bool){
    ie,ei := "ie","ei"
    var x int = 0
    v = false
        if(strings.Contains(s, ie)){
            x = strings.Index(s,ie)
            if !(s[x-1] == 'c'){
                v = true
            }
        }else if (strings.Contains(s, ei)){
            x = strings.Index(s,ei)
            if (s[x-1] == 'c'){
                v = true
        }
        }else{
            v = true
        }
    return v;
}

func check2(s string)bool{
    if(strings.Contains(s,"cie")){
        return false
    }
       if(strings.Contains(s,"ei")){
           if(!strings.Contains(s,"cei")){
           return false
           }
       }
       return true
}
