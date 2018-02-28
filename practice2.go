package main

import (
	"fmt"
)
type interf interface{
	interface_function()
}

type interf_example struct{
	str string
	number int
}

type Dog struct{
	name,breed string
	age, size int
}

type Account struct{
	id int
	balance float64
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


func x(ty int) func(num int) int{ 
	y := 10
	return func(num2 int) int{
		return (y * num2)/ty
	}
}

func dx(f1 func(int, int) int) int{ 			//taking function as parameter
	num1,num2 := 10,20
	return f1(num1,num2)
}

func (d Dog) printInfo() {
	fmt.Println("Dog",d.name,"is a",d.breed,"of",d.age,"years and weighs",d.size,"pounds")
}

func (a *Account) addMoney(money float64){
	a.balance += money
}

func addMoney2(a *Account, money float64){		//this is the same as above, its just a function with no receiver
	a.balance += money
} 

func (a Account) printBalance(){
	fmt.Println("User #",a.id,"Your balance is",a.balance)
}


func (i interf_example) interface_function(){
	fmt.Println(i.str,i.number)
}
func main() {
	t := x(5)			//closure example func returns a func using outside data
	fmt.Println("The number is: ",t(6)) 
	
	f := func(c,b int) int{		//assigning a function to a variable (like in javascript) so it can be used as a parameter
		return (c + b) * 100
	}
	
	a_num := dx(f)
	fmt.Println("The function is: ",a_num)
	
	d := Dog{"Ruff","Shiba Inu",5,25}
	d.printInfo()
	
	a := Account{0,10000}		//function using pointers with a receiver
	a.printBalance()
	a.addMoney(2000)
	a.printBalance()
	addMoney2(&a,5000)		//function just using pointers
	a.printBalance()
	
	p := &Account{1,90000}		//ya esta apuntando a una cuenta
	addMoney2(p,1000)		//notas que ya no necesitamos el ampersand?
	p.printBalance()
	
	
	var i interf = interf_example{"Hello",100}
	i.interface_function()
	
	do(21)
	do("Hello")
	
	

}
