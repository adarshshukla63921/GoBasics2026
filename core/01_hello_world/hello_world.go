// A package is just a collection of go files in the same folder with the same namespace
// Go wants explicit package declaration


// The entry of our program must be package 'main' , you can have other packages but not run them directly.
// other packages will be imported as part of main, if required

// The file name does not matter here, so we can just name it 01_hello_world.go , and that should be good.
// it is however conventional to see a 'main.go' file in projects, to indicate that this is the starting point.

package main

import "fmt"

// this is the entry point of our program.
func main() {
	fmt.Println("Hello, world")
}