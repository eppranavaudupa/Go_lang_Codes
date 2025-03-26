package main

import "fmt"

func main() {
	var a int
	println("Enter 1st number")
	fmt.Scan(&a)
	var b int
	println("Enter 2nd number")
	fmt.Scan(&b)
	var c int
	println("ENTER 3Rd number")
	fmt.Scan(&c)
	if a > b && a > c {
		println("a\t", a)
	} else if b > c {
		println("b", b)

	} else {
		println("c", c)
	}

}
