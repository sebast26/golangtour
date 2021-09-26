package main

import "fmt"

var num = []int{0,1}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	return func() int {
		v := num[len(num) - 1] + num[len(num) - 2]
		num = append(num, v)
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
