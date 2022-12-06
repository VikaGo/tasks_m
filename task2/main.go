package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	a, b, err := MiniMaxSum([]int{1, 2, 3, 4, 5})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
	fmt.Println(b)
}

func MiniMaxSum(arr []int) (int, int, error) {
	var _arr []int
	var i, j int

	if len(arr) == 0 {
		return 0, 0, errors.New("arr cant be nil")
	}

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
	return _arr[0], _arr[arrayLength-1], nil

}
