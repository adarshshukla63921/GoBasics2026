package main

import (
	"fmt"
	mathutil "package_in_go/math_util" // mathutil is an alias for the package.
	stringpackage "package_in_go/string_package" // stringpackage is an alias for the package
)

// what are packages in go ?

// they essentially what files / pieces of code belong to a single unit. helps keep them organized.
// we are familiar with main package already.

// the main package must have a main fucntion, which is the entry point of the program

// other packages, are imported in the main package and then used. you cannot use them directly.

// lets how to create our own package.

// first run : god mod init pacakge_name.
// create sub folder with name of package and create a go file with desired code.
// then we can import that package in main.go file
// import path : import "package_name/sub_folder_name"
// you can have only 1 package per directory, go enforeces this rule.

func main(){
	sum := mathutil.Add(2,4)
	fmt.Println(sum)

	sub := mathutil.Subtract(4,2)
	fmt.Println(sub)

	str := stringpackage.AddString("Adarsh ","Shukla")
	fmt.Println(str)
}