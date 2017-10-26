package main

import "fmt"

// func zero(x int) {
// 	x = 0
// }

//use pointer
func zero(xPtr *int) {
	*xPtr = 0
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	// x := 5
	// zero(&x)
	//x is unmodified by zero without pointer
	// fmt.Println(x)
	x := 1
	y := 2
	// square(&x)
	swap(&x, &y)
	fmt.Println(x, y)
}
