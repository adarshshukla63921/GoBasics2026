package main

import (
	"fmt"
	"math/cmplx"
)

// you can factor them like this , as in import statements

// Note that go values are initalized with 0 values. 
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<63 - 1
	z      complex128 = cmplx.Sqrt(-25 + 1i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// type conversion in go
	// go requires explicit conversions, unlike C or Java.
	var someInteger int = 43
	var someFloat float64 = float64(someInteger)
	fmt.Println(someFloat)

	// go has type inference, it can deduce the type of the variable based upon the value assigned to it.

	i := 42 // is an int
	j := 3.14257 // is a float64
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", j, j)
}