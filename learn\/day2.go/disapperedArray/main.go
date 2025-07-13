package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	sort.Ints(arr)
	fmt.Println(arr)

}
