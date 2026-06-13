package main

import "fmt"

// Functions are just reusable blocks of code, called with a name.

// this function takes 2 int and then returns their sum [ which is also an int ]
// a func may or may not accpet params or return one or many values.
// you can omit the type of params if they share the same type, for all but the last one
// func( a int ,b int) == func(a, b int)
func add(a, b int) int{
	return a + b
}

// func with multiple return values
func swap(a, b string) (string, string){
	return b,a
}

// func with named return values
// you should only use this in small functions, for better readability.
func swapDouble(a,b float32) (x,y float32){
	x=b
	y=a
	return
}

func greeting(){
	fmt.Println("Hello, welcome to using functions in go.")
}

func main(){
	greeting()
	// this is how you call a function
	fmt.Println("Sum of 2 and 3 is : ",add(2,3))

	// you could also do this : 
	x := add(5,7)
	println("sum of 5 and 7 is : ",x)

	// using multiple returns
	p,q := swap("world","hello")
	fmt.Println("Swapped values are : ",p,q)

	u,v := swapDouble(3.1,6.5)
	fmt.Println("Swapped values are : ",u,v)

}