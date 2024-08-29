package main

import (
	"fmt"
)

func main() {
	var longSide int = 5
	var shortSide int = 4
	var area = FindRectangleArea(5, 4)
	var perimeter = FindRectanglePerimeter(5, 4)
	fmt.Printf("Area of rectangle with side: %d and side: %d is %d\n", longSide, shortSide, area)
	fmt.Printf("Perimeter of rectangle with side: %d and side %d is %d\n", longSide, shortSide, perimeter)

	var str string = "dang"
	fmt.Printf("Response true is expected: %t\n", FindStringLengthBoolean(str))
	str = "hai"
	fmt.Printf("Response false is expected: %t\n", FindStringLengthBoolean(str))

	list := []int{4, 2, 5, 8, 1, 3}
	min, max, avg := FindSliceMinMaxAvg(list)
	fmt.Printf("Min, max, avg of a slice %v : %d, %d, %.2f\n", list, min, max, avg)

	list = []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Index of 2 nums such that they add up to %v is %v\n", target, TwoSum(list, target))
}
