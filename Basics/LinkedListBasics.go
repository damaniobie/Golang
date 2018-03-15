package main

import "fmt"
/* very basic intro to linked lists */
type Student struct{
    age int 
    weight int
    name string
    Next * Student
}

type Teacher struct{
    age int 
    weight int
    name string
    Next * Student
}
func main() {
    kyle := Student{7,100,"Kyle",nil}
    john := Student{10,70,"John",&kyle}
    
    deb := Teacher{24,140,"Deb",&john}
    
    fmt.Println("Teacher:",deb.name)
    fmt.Println("Behind john: ",deb.Next.Next.name)
    

}
