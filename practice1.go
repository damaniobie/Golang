
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
	
	for i,j := 0,0;i<5;i, j = i+1,j+2{	//initialize two vars in a for loop
		fmt.Println(i)	
	}

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

