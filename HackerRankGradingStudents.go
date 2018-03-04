package main
import "fmt"

/*
Sam is a professor at the university and likes to round each student's grade according to these rules:
If the difference between the grade and the next multiple of 5 is less than 3, round grade up to the next multiple of 5.
If the value of grade is less than 38, no rounding occurs as the result will still be a failing grade.
For example, grade = 84 will be rounded to 85 but grade = 29 will not be rounded because the rounding would result in a number 
that is less than 40.

Given the initial value of grade for each of Sam's n students, write code to automate the rounding process. Complete the function solve
that takes an integer array of all grades, and return an integer array consisting of the rounded grades. For each grade[i], round it
according to the rules above and print the result on a new line.
*/

var size int = 4
func main() {
    var x =[]int{73,67,38,33} //tested with various test cases
    solve(x)
    print_arr(x)
}

func solve(grades []int) []int{
    y := make([]int,len(grades))
    for i:=0;i<len(grades);i++{
        if(grades[i] >= 38){
            next_five := next5(grades[i])       
            if(next_five - grades[i] < 3){
                grades[i] = next_five
                y[i] = grades[i];
            }else{
                y[i] = grades[i];
            }
        }else{
            y[i] = grades[i];
        }
    }
    return y
    
}

func next5(x int)int{ //finding the next multiple of 5
    n := x
    if(x % 10==1 || x % 10 == 6){
        n = x+4
    }else if(x % 10==2 || x % 10 == 7){
        n = x+3
    }else if(x % 10==3 || x % 10 == 8){
        n = x+2
    }else if(x % 10==4 || x % 10 == 9){
        n = x+1
    }else{
    }
    return n
}

func print_arr(arr []int){
    for i := 0;i<len(arr);i++{
        fmt.Println(arr[i])
    }
}
