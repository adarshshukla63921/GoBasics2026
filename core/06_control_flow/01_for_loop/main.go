package main

import "fmt"

// go has only 1 looping construct, which is for loop.

func main(){
	// sum of first 10 natural numbers.
	sum := 0
	for i:=1; i<=10; i++ {
		sum += i
	}
	fmt.Println("Sum of first 10 natural number is : ",sum)

	// 'while' in for.
	redSum := 0
	for ; redSum < 100; {
		redSum += 10
	}
	
	fmt.Println("Reduced sum : ",redSum)

	// foever is for {...} , but do not write infinte loops.
}