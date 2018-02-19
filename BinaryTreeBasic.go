package main

import ("fmt"
        "math/rand"
       )
       

type node struct{
    left *node
    right *node
    val int
}

func insert(head *node,val int){
    newNode :=new(node)
    newNode.val = val
    temp := head
    
    for{
        if val < temp.val {
            if temp.left == nil{
                temp.left = newNode
                return
            }
            temp = temp.left
        }else if val >= temp.val{
            if temp.right == nil{
                temp.right = newNode
                return
            }
            temp = temp.right
        }
    }
}

func main() {
/* TEST */
    var head *node
    head = new(node)
    head.val = 50 //make this a higher number to assure we can get some values on the left side
    insert(head, 52)
    insert(head, 52) //inserting some that'll go right
    
   // fmt.Println(head.val)
   // fmt.Println(head.right.val)
    
    for i:= 0;i<50;i++{
        r := rand.Intn(100)
        insert(head,r)
    }
   
    fmt.Println(head.right.left.val)

}
