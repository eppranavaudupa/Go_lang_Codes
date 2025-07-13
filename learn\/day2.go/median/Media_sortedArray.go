package main

import "fmt"

func median(data []int) float64 {
	var length int = len(data)
	if length%2 != 0 {
		med := length / 2
		// fmt.Println(med, "med")
		ele := data[med]
		// fmt.Println(ele, "ELE")
		return float64(ele)
	}
	mid1 := data[length/2-1]
	mid2 := data[length/2]
	return float64(mid1+mid2) / 2
}
func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(median(arr))

}
