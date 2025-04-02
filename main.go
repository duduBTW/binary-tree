package main

import "fmt"

var size = 100000
var array = make([]int, size)

func find(target int) bool {
	l := 0
	r := size - 1

	for l <= r {
		m := (l + r) / 2
		value := array[m]

		if target > value {
			l = m + 1
		} else if target < value {
			r = m - 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Creating array...")

	for i := 0; i < size; i++ {
		array[i] = i
	}
	fmt.Println("Array created!")

	fmt.Println(find(size))
}
