package main

import "fmt"

func mul(a, b int) int {
	return a * b
}
func add(a, b int) (res int) {
	res = a + b
	return res
}
func main() {
	adder := add(2, 3) + mul(2, 5)
	fmt.Println(adder)

}
