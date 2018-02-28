The Go PlaygroundRun  Format   Imports  Share About
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50

package main

import (
	"fmt"
)

func main() {
	arr := make([]int,10)
	for i := range arr{
	arr[i] = i*5
	}
	fmt.Println(arr)

	/*kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
}*/
	//size := 100
	//arr2 := make([]int,size)
	var arr3 [8][8]int
	for i := 0; i<8;i++{
	for j := 0;j<8;j++{
		arr3[i][j] = i*j
	}
	}
	for i:=0;i<8;i++{
	for j:=0;j<8;j++{
    	fmt.Printf("arr3[%d][%d] = %d\n",i,j,arr3[i][j])
	}
	}
	
	//bubble sort
	for i:=1;i<len(arr);i++{
		for j:=0;j<len(arr)-1;j++{
			if(arr[j+1] > arr[j]){
				arr[j+1],arr[j] = arr[j],arr[j+1]
			}
		}
	}
		fmt.Println(arr)
	/* for  i := 0; i < 5; i++ {
      for j := 0; j < 2; j++ {
         fmt.Printf("arr3[%d][%d] = %d\n", i,j, arr3[i][j] )
      }
   }*/
	g := 100
	fmt.Printf("g is %d",g)

}

