package main

import (
	"example/pkg/genny"
	"fmt"
)

func main() {
	num := []int{1, 3, 5, 7, 9, 11}
	num = genny.Filter(num, func(element int) bool {
		if element < 6 {
			return true
		}
		return false
	})
	fmt.Println(num)
}
