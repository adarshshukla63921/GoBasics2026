package mathutil

// this is our dummy package.

// note exported functions are to be capitalized.
// this is how visibility works in go

// a name is exported if it begins with a capital letter.
// Any unexported names are not accessbile from other packages.
func Add(a , b int) int {
	return a + b
}

func Subtract(a , b int) int {
	return a - b
}