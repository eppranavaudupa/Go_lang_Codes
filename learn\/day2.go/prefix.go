package main

import (
	"fmt"
)

func pre(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	res := ""
	for i := 0; i < len(arr[0]); i++ {
		char := arr[0][i]
		for j := range arr {
			if i >= len(arr[j]) || arr[j][i] != char {
				return res
			}
		}
		res = res + string(char)
	}
	return res
}
func main() {
	data := []string{"flower", "flow", "flight"}
	fmt.Println(pre(data))

}
