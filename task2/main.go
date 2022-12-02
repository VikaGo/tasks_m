package main

import (
	"fmt"
	"sort"
)

func MiniMaxSum(arr []int) (int, int) {
	var _arr []int
	var i, j int

	arrayLength := len(arr)

	for i = 0; i < arrayLength; i++ {
		var sum int
		for j = 0; j < arrayLength; j++ {
			if i != j {
				sum += arr[j]
			}
		}
		_arr = append(_arr, sum)

	}

	sort.Ints(_arr)
	return _arr[0], _arr[arrayLength-1]

}

func main() {
	a, b := MiniMaxSum([]int{1, 2, 3, 4, 5})
	fmt.Println(a)
	fmt.Println(b)

}
