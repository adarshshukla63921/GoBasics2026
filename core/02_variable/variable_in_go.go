package main

// This file will demonstrate on how to use go variables.
// Inside go, you must use the variables, import etc , else you will get an error.

var appName = "learning go" // only const, var is allowed here and not := 

func main(){
	// var is suited for package level declarations or when explicit type is needed.

	var x int
	println(x) // go variables are initialized with zero value.
	x = 10 // assigning a new value to an existing variable.
	println(x) 
	z := 20 // short variable declaration - type is inferred from assigned value
	println(z)

	// inside function ' := ' is the most preferred form of variable declaration.
	// x := 20 is not allowed, as x already exists in the same scope.

	newVar := "Hello, Go!"
	println(newVar) // some new variable can use := in the same scope.

	// grouped delcarations also use var, good for readablity.
	var(
		port = 8080
		host = "localhost"
	)

	println("running server at", host, "on port", port)


	// you can also use constants in go

	const fixedValue = "i will learn go"
	// fixedValue = "not now" will cause error as constants cannot be reassgined.



}