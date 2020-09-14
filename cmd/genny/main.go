package main

import (
	"example/pkg/genny"
	"fmt"
)

func main() {
	num := []int{1, 3, 5, 7, 9, 11}
	num2 := []float32{1.0, 3.0, 5.0, 7.0, 9.0, 11.0}
	var out interface{}
	// filter by static filter
	num = genny.Filter(num, func(element int) bool {
		if element < 6 {
			return true
		}
		return false
	})
	fmt.Println(num)
	// filter by int filter
	out = genny.FilterWithInterface(num, func(element interface{}) bool {
		if element.(int) < 6 {
			return true
		}
		return false
	})
	fmt.Println(out.([]int))
	// filter by float filter
	out = genny.FilterWithInterface(num2, func(element interface{}) bool {
		if element.(float32) < 6 {
			return true
		}
		return false
	})
	fmt.Println(out.([]float32))
}
